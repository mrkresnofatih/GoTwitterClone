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

	getHomeFeedEndpoint := &utils.ApplicationEndpoint{
		Handler: timelinehandlers.GetHomeFeedHandler,
		Path:    "/get-home",
		Method:  http.MethodPost,
	}
	getHomeFeedWithValidation := &utils.RequireValidation[models.HomeFeedQueryModel]{
		Endpoint: getHomeFeedEndpoint,
	}
	getHomeFeedWithAuth := &utils.RequireAuthentication{
		Endpoint: getHomeFeedWithValidation,
	}
	subRouter.AddEndpoint(getHomeFeedWithAuth)

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

	getRepliesFeedEndpoint := &utils.ApplicationEndpoint{
		Handler: timelinehandlers.GetReplyFeedHandler,
		Path:    "/get-replies",
		Method:  http.MethodPost,
	}
	getRepliesFeedWithValidation := &utils.RequireValidation[models.ReplyFeedQueryModel]{
		Endpoint: getRepliesFeedEndpoint,
	}
	getRepliesFeedWithAuth := &utils.RequireAuthentication{
		Endpoint: getRepliesFeedWithValidation,
	}
	subRouter.AddEndpoint(getRepliesFeedWithAuth)

	subRouter.Init()
}
