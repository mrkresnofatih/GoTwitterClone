package followhandlers

import (
	"encoding/json"
	"log"
	"mrkresnofatihdev/apps/gotwittercloneapp/application/followservice"
	"mrkresnofatihdev/apps/gotwittercloneapp/models"
	"mrkresnofatihdev/apps/gotwittercloneapp/utils"
	"net/http"
)

func GetFollowerListProfilesHandler(w http.ResponseWriter, r *http.Request) {
	responseHelper := utils.ResponseHelper{Writer: w}
	var getFollowersListQuery models.FollowListQueryModel
	_ = json.NewDecoder(r.Body).Decode(&getFollowersListQuery)

	followerProfiles, err := followservice.GetFollowerListProfiles(r.Context(), getFollowersListQuery)
	if err != nil {
		log.Println("get-follower-list-profiles failed")
		responseHelper.SetJsonResponse(http.StatusBadRequest, []string{
			"get-follower-list-profiles failed",
		})
		return
	}
	responseHelper.SetJsonResponse(http.StatusOK, followerProfiles)
}
