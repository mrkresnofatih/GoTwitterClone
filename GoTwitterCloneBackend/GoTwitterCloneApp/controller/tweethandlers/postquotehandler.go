package tweethandlers

import (
	"encoding/json"
	"mrkresnofatihdev/apps/gotwittercloneapp/application/tweetservice"
	"mrkresnofatihdev/apps/gotwittercloneapp/models"
	"mrkresnofatihdev/apps/gotwittercloneapp/utils"
	"net/http"
)

func PostQuoteHandler(w http.ResponseWriter, r *http.Request) {
	responseHelper := utils.ResponseHelper{Writer: w}
	var postQuoteReq models.TweetReplyRequestModel
	_ = json.NewDecoder(r.Body).Decode(&postQuoteReq)

	authHeader := r.Header.Get(utils.HttpHeaderKeyAuthorization)
	token := authHeader[7:]
	requesterUsername, err := utils.GetClaimFromToken[string](token, utils.ApplicationJwtClaimsKeyUsername)
	if err != nil {
		responseHelper.SetJsonResponse(http.StatusBadRequest, []string{
			"username_invalid",
		})
		return
	}

	newlyCreatedTweet, err := tweetservice.PostQuote(r.Context(), postQuoteReq, requesterUsername)
	if err != nil {
		responseHelper.SetJsonResponse(http.StatusBadRequest, []string{
			"error_create_tweet",
		})
		return
	}

	responseHelper.SetJsonResponse(http.StatusOK, newlyCreatedTweet)
}