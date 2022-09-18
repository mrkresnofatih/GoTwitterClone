package timelinehandlers

import (
	"encoding/json"
	"log"
	"mrkresnofatihdev/apps/gotwittercloneapp/application/timelineservice"
	"mrkresnofatihdev/apps/gotwittercloneapp/models"
	"mrkresnofatihdev/apps/gotwittercloneapp/utils"
	"net/http"
)

func GetReplyFeedHandler(w http.ResponseWriter, r *http.Request) {
	responseHelper := utils.ResponseHelper{Writer: w}
	var getReplies models.ReplyFeedQueryModel
	_ = json.NewDecoder(r.Body).Decode(&getReplies)

	replyFeedResponseData, err := timelineservice.GetReplyFeed(r.Context(), getReplies)
	if err != nil {
		log.Println("error: " + err.Error())
		responseHelper.SetJsonResponse(http.StatusBadRequest, []string{
			err.Error(),
		})
		return
	}

	responseHelper.SetJsonResponse(http.StatusOK, replyFeedResponseData)
}
