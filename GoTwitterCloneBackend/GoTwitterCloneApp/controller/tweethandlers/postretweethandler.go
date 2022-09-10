package tweethandlers

import (
	"encoding/json"
	"mrkresnofatihdev/apps/gotwittercloneapp/application/tweetservice"
	"mrkresnofatihdev/apps/gotwittercloneapp/events"
	"mrkresnofatihdev/apps/gotwittercloneapp/models"
	"mrkresnofatihdev/apps/gotwittercloneapp/utils"
	"net/http"
)

func PostRetweetHandler(w http.ResponseWriter, r *http.Request) {
	responseHelper := utils.ResponseHelper{Writer: w}
	var postRetweetReq models.TweetRetweetRequestModel
	_ = json.NewDecoder(r.Body).Decode(&postRetweetReq)

	authHeader := r.Header.Get(utils.HttpHeaderKeyAuthorization)
	token := authHeader[7:]
	requesterUsername, err := utils.GetClaimFromToken[string](token, utils.ApplicationJwtClaimsKeyUsername)
	if err != nil {
		responseHelper.SetJsonResponse(http.StatusBadRequest, []string{
			"username_invalid",
		})
		return
	}

	newlyCreatedTweet, err := tweetservice.PostRetweet(r.Context(), postRetweetReq, requesterUsername)
	if err != nil {
		responseHelper.SetJsonResponse(http.StatusBadRequest, []string{
			"error_create_tweet",
		})
		return
	}

	events.PublishEventMessage(r.Context(), events.IncrementRetweetCountEventHandlerName, postRetweetReq)
	responseHelper.SetJsonResponse(http.StatusOK, newlyCreatedTweet)
}
