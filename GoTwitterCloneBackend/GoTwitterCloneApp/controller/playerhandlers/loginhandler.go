package playerhandlers

import (
	"encoding/json"
	"mrkresnofatihdev/apps/gotwittercloneapp/application/playerservice"
	"mrkresnofatihdev/apps/gotwittercloneapp/models"
	"mrkresnofatihdev/apps/gotwittercloneapp/utils"
	"net/http"
)

func LoginPlayerHandler(w http.ResponseWriter, r *http.Request) {
	responseHelper := utils.ResponseHelper{Writer: w}
	var loginReq models.PlayerLoginRequestModel
	err := json.NewDecoder(r.Body).Decode(&loginReq)
	if err != nil {
		responseHelper.SetJsonResponse(http.StatusBadRequest, []string{
			"json_body_parse_failed",
		})
		return
	}

	loginResp, err := playerservice.LoginPlayer(r.Context(), loginReq)
	if err != nil {
		responseHelper.SetJsonResponse(http.StatusBadRequest, []string{
			"login_failed",
		})
		return
	}
	responseHelper.SetJsonResponse(http.StatusOK, loginResp)

}
