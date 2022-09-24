package timelinehandlers

import (
	"encoding/json"
	"mrkresnofatihdev/apps/gotwittercloneapp/application/timelineservice"
	"mrkresnofatihdev/apps/gotwittercloneapp/events"
	"mrkresnofatihdev/apps/gotwittercloneapp/models"
	"mrkresnofatihdev/apps/gotwittercloneapp/utils"
	"net/http"
)

func GetHomeFeedHandler(w http.ResponseWriter, r *http.Request) {
	responseHelper := utils.ResponseHelper{Writer: w}
	var getHomeQuery models.HomeFeedQueryModel
	_ = json.NewDecoder(r.Body).Decode(&getHomeQuery)
	authHeader := r.Header.Get(utils.HttpHeaderKeyAuthorization)
	token := authHeader[7:]

	requesterUsername, err := utils.GetClaimFromToken[string](token, utils.ApplicationJwtClaimsKeyUsername)
	if err != nil {
		responseHelper.SetJsonResponse(http.StatusBadRequest, []string{
			"username_doesnt_exist_in_token_claims",
		})
		return
	}

	homeFeedRunning, err := timelineservice.GetHomeFeedQueryRunning(r.Context(), requesterUsername)
	if err != nil {
		responseHelper.SetJsonResponse(http.StatusBadRequest, []string{
			"failed when get-home-feed-query-running",
		})
		return
	}

	if !homeFeedRunning {
		events.PublishEventMessage(r.Context(), events.RunHomeFeedQueryEventHandlerName, requesterUsername)
	}

	homeFeedBatchTweets, err := timelineservice.GetHomeFeed(r.Context(), getHomeQuery, requesterUsername)
	if err != nil {
		responseHelper.SetJsonResponse(http.StatusBadRequest, []string{
			"get_feed_batch_tweets_failed",
		})
		return
	}

	responseHelper.SetJsonResponse(http.StatusOK, homeFeedBatchTweets)
}
