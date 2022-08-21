package playerhandlers

import (
	"encoding/json"
	"mrkresnofatihdev/apps/gotwittercloneapp/application/playerservice"
	"mrkresnofatihdev/apps/gotwittercloneapp/models"
	"mrkresnofatihdev/apps/gotwittercloneapp/utils"
	"net/http"
)

func ValidateCreatePlayerHandler(w http.ResponseWriter, r *http.Request) {
	responseHelper := utils.ResponseHelper{Writer: w}
	var createPlayerReq models.PlayerCreateRequestModel
	err := json.NewDecoder(r.Body).Decode(&createPlayerReq)
	if err != nil {
		responseHelper.SetJsonResponse(http.StatusBadRequest, []string{
			err.Error(),
		})
		return
	}

	errorList := playerservice.ValidateCreatePlayer(r.Context(), createPlayerReq)

	isPlayerExists := playerservice.GetPlayerExists(r.Context(), createPlayerReq.Username)
	if isPlayerExists {
		errorList = append(errorList, "player_username_exists")
	}
	responseHelper.SetJsonResponse(http.StatusOK, errorList)
}
