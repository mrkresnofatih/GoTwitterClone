package followhandlers

import (
	"github.com/gorilla/mux"
	"mrkresnofatihdev/apps/gotwittercloneapp/application/followservice"
	"mrkresnofatihdev/apps/gotwittercloneapp/models"
	"mrkresnofatihdev/apps/gotwittercloneapp/utils"
	"net/http"
)

func StartFollowHandler(w http.ResponseWriter, r *http.Request) {
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
	targetFollowUsername := pathVars["username"]

	startFollowRequest := models.FollowRequestModel{
		Username:         targetFollowUsername,
		FollowerUsername: requesterUsername,
	}

	err = followservice.StartFollowing(r.Context(), startFollowRequest)
	if err != nil {
		responseHelper.SetJsonResponse(http.StatusBadRequest, []string{
			"start_follow_attempt_failed",
		})
		return
	}

	responseHelper.SetJsonResponse(http.StatusOK, nil)
}
