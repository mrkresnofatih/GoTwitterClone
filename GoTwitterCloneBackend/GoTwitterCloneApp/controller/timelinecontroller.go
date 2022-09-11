package controller

import (
	"github.com/gorilla/mux"
	"mrkresnofatihdev/apps/gotwittercloneapp/controller/timelinehandlers"
	"mrkresnofatihdev/apps/gotwittercloneapp/models"
	"mrkresnofatihdev/apps/gotwittercloneapp/utils"
	"net/http"
)

type TimelineController struct{}

func (t *TimelineController) AddControllerTo(router *mux.Router) {
	subRouter := &utils.ApplicationRouter{
		Parent:     router,
		PathPrefix: "/timeline",
	}

	getProfileFeedEndpoint := &utils.ApplicationEndpoint{
		Handler: timelinehandlers.GetProfileFeedHandler,
		Path:    "/get-profile",
		Method:  http.MethodPost,
	}
	getProfileFeedWithValidation := &utils.RequireValidation[models.ProfileFeedQueryModel]{
		Endpoint: getProfileFeedEndpoint,
	}
	getProfileFeedWithAuth := &utils.RequireAuthentication{
		Endpoint: getProfileFeedWithValidation,
	}
	subRouter.AddEndpoint(getProfileFeedWithAuth)

	subRouter.Init()
}
