package utils

import (
	"github.com/gorilla/mux"
	"net/http"
)

type IServer interface {
	Initialize()
	AddController(IController)
}

type ApplicationServer struct {
	MainRouter  *mux.Router
	Controllers []IController
}

func (s *ApplicationServer) Initialize() {
	s.MainRouter = mux.NewRouter()
	for _, controller := range s.Controllers {
		controller.AddControllerTo(s.MainRouter)
	}
}

func (s *ApplicationServer) AddController(g IController) {
	s.Controllers = append(s.Controllers, g)
}

type IController interface {
	AddControllerTo(*mux.Router)
}

type IRouter interface {
	Init()
	AddEndpoint(IEndpoint)
}

type ApplicationRouter struct {
	Parent     *mux.Router
	Router     *mux.Router
	PathPrefix string
	Endpoints  []IEndpoint
}

func (s *ApplicationRouter) Init() {
	s.Router = s.Parent.PathPrefix(s.PathPrefix).Subrouter()
	for _, endpoint := range s.Endpoints {
		endpoint.AddEndpointTo(s.Router)
	}
}

func (s *ApplicationRouter) AddEndpoint(i IEndpoint) {
	s.Endpoints = append(s.Endpoints, i)
}

type HandlerFunc func(http.ResponseWriter, *http.Request)

type IEndpoint interface {
	GetMethod() string
	GetPath() string
	GetHandler() HandlerFunc
	AddEndpointTo(*mux.Router)
}

type ApplicationEndpoint struct {
	Handler HandlerFunc
	Method  string
	Path    string
}

func (s *ApplicationEndpoint) GetHandler() HandlerFunc {
	return s.Handler
}

func (s *ApplicationEndpoint) GetMethod() string {
	return s.Method
}

func (s *ApplicationEndpoint) GetPath() string {
	return s.Path
}

func (s *ApplicationEndpoint) AddEndpointTo(router *mux.Router) {
	router.
		HandleFunc(s.GetPath(), s.GetHandler()).
		Methods(s.GetMethod())
}
