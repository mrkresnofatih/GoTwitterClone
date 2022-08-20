package playerhandlers

import (
	"encoding/json"
	"mrkresnofatihdev/apps/gotwittercloneapp/application/playerservice"
	"mrkresnofatihdev/apps/gotwittercloneapp/models"
	"mrkresnofatihdev/apps/gotwittercloneapp/utils"
	"net/http"
)

func CreatePlayerHandler(w http.ResponseWriter, r *http.Request) {
	responseHelper := utils.ResponseHelper{Writer: w}
	var createPlayerRequest models.PlayerCreateRequestModel
	_ = json.NewDecoder(r.Body).Decode(&createPlayerRequest)

	playerCreateResponse, err := playerservice.CreatePlayer(r.Context(), createPlayerRequest)
	if err != nil {
		responseHelper.SetJsonResponse(http.StatusBadRequest, []string{
			err.Error(),
		})
		return
	}
	responseHelper.SetJsonResponse(http.StatusOK, playerCreateResponse)
}