package timelinehandlers

import (
	"encoding/json"
	"log"
	"mrkresnofatihdev/apps/gotwittercloneapp/application/timelineservice"
	"mrkresnofatihdev/apps/gotwittercloneapp/models"
	"mrkresnofatihdev/apps/gotwittercloneapp/utils"
	"net/http"
)

func GetProfileFeedHandler(w http.ResponseWriter, r *http.Request) {
	responseHelper := utils.ResponseHelper{Writer: w}
	var getProfile models.ProfileFeedQueryModel
	_ = json.NewDecoder(r.Body).Decode(&getProfile)

	log.Println("test")

	profileFeedResponseData, err := timelineservice.GetProfileFeed(r.Context(), getProfile)
	if err != nil {
		log.Println("error: " + err.Error())
		responseHelper.SetJsonResponse(http.StatusBadRequest, []string{
			err.Error(),
		})
		return
	}

	responseHelper.SetJsonResponse(http.StatusOK, profileFeedResponseData)
}