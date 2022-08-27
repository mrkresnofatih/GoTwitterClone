package controller

import (
	"github.com/gorilla/mux"
	"mrkresnofatihdev/apps/gotwittercloneapp/controller/followhandlers"
	"mrkresnofatihdev/apps/gotwittercloneapp/utils"
	"net/http"
)

type FollowController struct {}

func (f *FollowController) AddControllerTo(router *mux.Router) {
	subRouter := &utils.ApplicationRouter{
		Parent: router,
		PathPrefix: "/follow",
	}

	startFollowingEndpoint := &utils.ApplicationEndpoint{
		Handler: followhandlers.StartFollowHandler,
		Path: "/start/{username}",
		Method: http.MethodGet,
	}
	startFollowWithAuthentication := &utils.RequireAuthentication{
		Endpoint: startFollowingEndpoint,
	}
	subRouter.AddEndpoint(startFollowWithAuthentication)

	stopFollowingEndpoint := &utils.ApplicationEndpoint{
		Handler: followhandlers.StopFollowHandler,
		Path: "/stop/{username}",
		Method: http.MethodGet,
	}
	stopFollowWithAuthentication := &utils.RequireAuthentication{
		Endpoint: stopFollowingEndpoint,
	}
	subRouter.AddEndpoint(stopFollowWithAuthentication)

	subRouter.Init()
}
