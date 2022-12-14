package tweethandlers

import (
	"encoding/json"
	"mrkresnofatihdev/apps/gotwittercloneapp/application/tweetservice"
	"mrkresnofatihdev/apps/gotwittercloneapp/events"
	"mrkresnofatihdev/apps/gotwittercloneapp/models"
	"mrkresnofatihdev/apps/gotwittercloneapp/utils"
	"net/http"
)

func PostReplyHandler(w http.ResponseWriter, r *http.Request) {
	responseHelper := utils.ResponseHelper{Writer: w}
	var postReplyReq models.TweetReplyRequestModel
	_ = json.NewDecoder(r.Body).Decode(&postReplyReq)

	authHeader := r.Header.Get(utils.HttpHeaderKeyAuthorization)
	token := authHeader[7:]
	requesterUsername, err := utils.GetClaimFromToken[string](token, utils.ApplicationJwtClaimsKeyUsername)
	if err != nil {
		responseHelper.SetJsonResponse(http.StatusBadRequest, []string{
			"username_invalid",
		})
		return
	}

	newlyCreatedReply, err := tweetservice.PostReply(r.Context(), postReplyReq, requesterUsername)
	if err != nil {
		responseHelper.SetJsonResponse(http.StatusBadRequest, []string{
			"error_post_reply",
		})
		return
	}

	events.PublishEventMessage(r.Context(), events.IncrementReplyCountEventHandlerName, postReplyReq)
	events.PublishEventMessage(r.Context(), events.RecordReplyTweetEventHandlerName, newlyCreatedReply)

	responseHelper.SetJsonResponse(http.StatusOK, newlyCreatedReply)
}
