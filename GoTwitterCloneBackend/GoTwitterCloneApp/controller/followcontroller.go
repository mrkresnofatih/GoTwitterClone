package controller

import (
	"github.com/gorilla/mux"
	"mrkresnofatihdev/apps/gotwittercloneapp/controller/followhandlers"
	"mrkresnofatihdev/apps/gotwittercloneapp/models"
	"mrkresnofatihdev/apps/gotwittercloneapp/utils"
	"net/http"
)

type FollowController struct{}

func (f *FollowController) AddControllerTo(router *mux.Router) {
	subRouter := &utils.ApplicationRouter{
		Parent:     router,
		PathPrefix: "/follow",
	}

	startFollowingEndpoint := &utils.ApplicationEndpoint{
		Handler: followhandlers.StartFollowHandler,
		Path:    "/start/{username}",
		Method:  http.MethodGet,
	}
	startFollowWithAuthentication := &utils.RequireAuthentication{
		Endpoint: startFollowingEndpoint,
	}
	subRouter.AddEndpoint(startFollowWithAuthentication)

	stopFollowingEndpoint := &utils.ApplicationEndpoint{
		Handler: followhandlers.StopFollowHandler,
		Path:    "/stop/{username}",
		Method:  http.MethodGet,
	}
	stopFollowWithAuthentication := &utils.RequireAuthentication{
		Endpoint: stopFollowingEndpoint,
	}
	subRouter.AddEndpoint(stopFollowWithAuthentication)

	getFollowerListProfilesEndpoint := &utils.ApplicationEndpoint{
		Handler: followhandlers.GetFollowerListProfilesHandler,
		Path:    "/get-followers",
		Method:  http.MethodPost,
	}
	getFollowerProfilesWithValidation := &utils.RequireValidation[models.FollowListQueryModel]{
		Endpoint: getFollowerListProfilesEndpoint,
	}
	getFollowerProfilesWithAuth := &utils.RequireAuthentication{
		Endpoint: getFollowerProfilesWithValidation,
	}
	subRouter.AddEndpoint(getFollowerProfilesWithAuth)

	getFollowingListProfilesEndpoint := &utils.ApplicationEndpoint{
		Handler: followhandlers.GetFollowingListProfilesHandler,
		Path:    "/get-followings",
		Method:  http.MethodPost,
	}
	getFollowingProfilesWithValidation := &utils.RequireValidation[models.FollowListQueryModel]{
		Endpoint: getFollowingListProfilesEndpoint,
	}
	getFollowingProfilesWithAuth := &utils.RequireAuthentication{
		Endpoint: getFollowingProfilesWithValidation,
	}
	subRouter.AddEndpoint(getFollowingProfilesWithAuth)

	subRouter.Init()
}
