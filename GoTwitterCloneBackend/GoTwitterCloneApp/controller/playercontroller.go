package controller

import (
	"github.com/gorilla/mux"
	"mrkresnofatihdev/apps/gotwittercloneapp/controller/playerhandlers"
	"mrkresnofatihdev/apps/gotwittercloneapp/models"
	"mrkresnofatihdev/apps/gotwittercloneapp/utils"
	"net/http"
)

type PlayerController struct{}

func (p *PlayerController) AddControllerTo(router *mux.Router) {
	subRouter := &utils.ApplicationRouter{
		Parent:     router,
		PathPrefix: "/player",
	}

	createPlayerEndpoint := &utils.ApplicationEndpoint{
		Handler: playerhandlers.CreatePlayerHandler,
		Path:    "/create",
		Method:  http.MethodPost,
	}
	createPlayerEndpointWithValidation := &utils.RequireValidation[models.PlayerCreateRequestModel]{
		Endpoint: createPlayerEndpoint,
	}
	subRouter.AddEndpoint(createPlayerEndpointWithValidation)

	validateCreatePlayerEndpoint := &utils.ApplicationEndpoint{
		Handler: playerhandlers.ValidateCreatePlayerHandler,
		Path: "/validate",
		Method: http.MethodPost,
	}
	subRouter.AddEndpoint(validateCreatePlayerEndpoint)

	loginPlayerEndpoint := &utils.ApplicationEndpoint{
		Handler: playerhandlers.LoginPlayerHandler,
		Path: "/login",
		Method: http.MethodPost,
	}
	loginPlayerWithValidation := &utils.RequireValidation[models.PlayerLoginRequestModel]{
		Endpoint: loginPlayerEndpoint,
	}
	subRouter.AddEndpoint(loginPlayerWithValidation)

	subRouter.Init()
}
