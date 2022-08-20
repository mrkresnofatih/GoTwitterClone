package utils

import (
	"bytes"
	"encoding/json"
	"fmt"
	validator2 "github.com/go-playground/validator/v10"
	"github.com/gorilla/mux"
	"io"
	"net/http"
)

//#region RequireAuthenticationDecorator
type RequireAuthentication struct {
	Endpoint IEndpoint
}

func (s *RequireAuthentication) GetHandler() HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		responseHelper := ResponseHelper{Writer: w}
		authorizationValue := r.Header.Get(HttpHeaderKeyAuthorization)
		jwtToken := authorizationValue[:7]
		isJwtTokenValid := GetValidityFromToken(jwtToken)
		if isJwtTokenValid {
			s.Endpoint.GetHandler()(w, r)
		} else {
			responseHelper.SetJsonResponse(http.StatusBadRequest, []string{
				"Auth_Failed",
			})
		}
	}
}

func (s *RequireAuthentication) GetPath() string {
	return s.Endpoint.GetPath()
}

func (s *RequireAuthentication) GetMethod() string {
	return s.Endpoint.GetMethod()
}

func (s *RequireAuthentication) AddEndpointTo(router *mux.Router) {
	router.
		HandleFunc(s.GetPath(), s.GetHandler()).
		Methods(s.GetMethod())
}
//#endregion

//#region RequirePermissionDecorator
type RequirePermission struct {
	Endpoint   IEndpoint
	Permission string
}

func (s *RequirePermission) GetHandler() HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		responseHelper := ResponseHelper{Writer: w}
		authorizationValue := r.Header.Get(HttpHeaderKeyAuthorization)
		jwtToken := authorizationValue[:7]
		permissionExists := GetPermissionExistsFromToken(jwtToken, s.Permission)
		if permissionExists {
			s.Endpoint.GetHandler()(w, r)
		} else {
			responseHelper.SetJsonResponse(http.StatusBadRequest, []string{
				"Invalid_Auth_Token",
			})
		}
	}
}

func (s *RequirePermission) GetPath() string {
	return s.Endpoint.GetPath()
}

func (s *RequirePermission) GetMethod() string {
	return s.Endpoint.GetMethod()
}

func (s *RequirePermission) AddEndpointTo(router *mux.Router) {
	router.
		HandleFunc(s.GetPath(), s.GetHandler()).
		Methods(s.GetMethod())
}
//#endregion

//#region RequireValidationDecorator
type RequireValidation[T interface{}] struct {
	Endpoint IEndpoint
}

func (s *RequireValidation[T]) GetPath() string {
	return s.Endpoint.GetPath()
}

func (s *RequireValidation[T]) GetMethod() string {
	return s.Endpoint.GetMethod()
}

func (s *RequireValidation[T]) GetHandler() HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		responseHelper := ResponseHelper{Writer: w}
		body, err := io.ReadAll(r.Body)
		var bodyData T
		err = json.Unmarshal(body, &bodyData)
		if err != nil {
			responseHelper.SetJsonResponse(http.StatusBadRequest, []string{
				"parse_json_failed",
			})
			return
		}

		validator := validator2.New()
		err = validator.Struct(bodyData)
		if err != nil {
			if _, ok := err.(*validator2.InvalidValidationError); ok {
				fmt.Println(err)
				return
			}

			for _, err := range err.(validator2.ValidationErrors) {
				fmt.Println(err)
			}

			responseHelper.SetJsonResponse(http.StatusBadRequest, []string{
				"request_body_validation_failed",
			})
			return
		}

		newR := r.Clone(r.Context())
		r.Body = io.NopCloser(bytes.NewReader(body))
		newR.Body = io.NopCloser(bytes.NewReader(body))
		err = r.ParseForm()
		if err != nil {
			responseHelper.SetJsonResponse(http.StatusBadRequest, []string{
				"request_validation_parse_form_failed",
			})
		}
		s.Endpoint.GetHandler()(w, newR)
	}
}

func (s *RequireValidation[T]) AddEndpointTo(router *mux.Router) {
	router.
		HandleFunc(s.GetPath(), s.GetHandler()).
		Methods(s.GetMethod())
}
//#endregion