package followhandlers

import (
	"encoding/json"
	"log"
	"mrkresnofatihdev/apps/gotwittercloneapp/application/followservice"
	"mrkresnofatihdev/apps/gotwittercloneapp/models"
	"mrkresnofatihdev/apps/gotwittercloneapp/utils"
	"net/http"
)

func GetFollowingListProfilesHandler(w http.ResponseWriter, r *http.Request) {
	responseHelper := utils.ResponseHelper{Writer: w}
	var getFollowingListQuery models.FollowListQueryModel
	_ = json.NewDecoder(r.Body).Decode(&getFollowingListQuery)

	followingProfiles, err := followservice.GetFollowingListProfiles(r.Context(), getFollowingListQuery)
	if err != nil {
		log.Println("get-following-list-profiles failed")
		responseHelper.SetJsonResponse(http.StatusBadRequest, []string{
			"get-following-list-profiles failed",
		})
		return
	}
	responseHelper.SetJsonResponse(http.StatusOK, followingProfiles)
}
