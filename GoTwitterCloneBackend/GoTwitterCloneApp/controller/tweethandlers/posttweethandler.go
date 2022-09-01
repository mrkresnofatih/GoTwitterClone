package tweethandlers

import (
	"encoding/json"
	"mrkresnofatihdev/apps/gotwittercloneapp/application/tweetservice"
	"mrkresnofatihdev/apps/gotwittercloneapp/models"
	"mrkresnofatihdev/apps/gotwittercloneapp/utils"
	"net/http"
)

func PostTweetHandler(w http.ResponseWriter, r *http.Request) {
	responseHelper := utils.ResponseHelper{Writer: w}
	var postTweetReq models.TweetCreateRequestModel
	_ = json.NewDecoder(r.Body).Decode(&postTweetReq)

	authHeader := r.Header.Get(utils.HttpHeaderKeyAuthorization)
	token := authHeader[7:]
	requesterUsername, err := utils.GetClaimFromToken[string](token, utils.ApplicationJwtClaimsKeyUsername)
	if err != nil {
		responseHelper.SetJsonResponse(http.StatusBadRequest, []string{
			"username_invalid",
		})
		return
	}

	newlyCreatedTweet, err := tweetservice.PostTweet(r.Context(), postTweetReq, requesterUsername)
	if err != nil {
		responseHelper.SetJsonResponse(http.StatusBadRequest, []string{
			"error_create_tweet",
		})
		return
	}

	responseHelper.SetJsonResponse(http.StatusOK, newlyCreatedTweet)
}