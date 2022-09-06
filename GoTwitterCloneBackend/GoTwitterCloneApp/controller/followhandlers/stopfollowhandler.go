package followhandlers

import (
	"github.com/gorilla/mux"
	"mrkresnofatihdev/apps/gotwittercloneapp/events"
	"mrkresnofatihdev/apps/gotwittercloneapp/events/eventhandlers"
	"mrkresnofatihdev/apps/gotwittercloneapp/models"
	"mrkresnofatihdev/apps/gotwittercloneapp/utils"
	"net/http"
)

func StopFollowHandler(w http.ResponseWriter, r *http.Request) {
	responseHelper := utils.ResponseHelper{Writer: w}
	authHeader := r.Header.Get(utils.HttpHeaderKeyAuthorization)
	token := authHeader[7:]

	requesterUsername, err := utils.GetClaimFromToken[string](token, utils.ApplicationJwtClaimsKeyUsername)
	if err != nil {
		responseHelper.SetJsonResponse(http.StatusBadRequest, []string{
			"username_doesnt_exist_in_token_claims",
		})
		return
	}

	pathVars := mux.Vars(r)
	targetUnfollowUsername := pathVars["username"]

	stopFollowRequest := models.FollowRequestModel{
		Username:         targetUnfollowUsername,
		FollowerUsername: requesterUsername,
	}

	events.PublishEventMessage(r.Context(), eventhandlers.UnfollowEventHandlerName, stopFollowRequest)
	responseHelper.SetJsonResponse(http.StatusOK, nil)
}
