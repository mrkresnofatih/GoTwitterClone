package playerservice

import (
	"context"
	"errors"
	"mrkresnofatihdev/apps/gotwittercloneapp/models"
	"mrkresnofatihdev/apps/gotwittercloneapp/utils"
	"time"
)

func LoginPlayer(ctx context.Context, loginReq models.PlayerLoginRequestModel) (models.PlayerLoginResponseModel, error) {
	returnDefaultsOnError := func(err error) (models.PlayerLoginResponseModel, error) {
		return *new(models.PlayerLoginResponseModel), err
	}

	hashUtil := utils.HashUtility{
		HashStrategy: &utils.HashUserPasswordStrategy{},
	}

	credential, err := GetPlayerCredentials(ctx, loginReq.Username)
	if err != nil {
		return returnDefaultsOnError(err)
	}

	passwordValid := hashUtil.GetIsCompareHashValid(credential.Password, loginReq.Password)
	if !passwordValid {
		return returnDefaultsOnError(errors.New("player_password_invalid"))
	}

	profile, err := GetPlayerMinimumProfile(ctx, loginReq.Username)
	if err != nil {
		return returnDefaultsOnError(err)
	}

	tokenBuilder := utils.BasicJwtTokenBuilder{
		ExpiresAfter: time.Hour * 1,
	}
	tokenWithUsrBuilder := utils.UsernameJwtTokenBuilder{
		Username:        loginReq.Username,
		JwtTokenBuilder: &tokenBuilder,
	}
	token, _ := tokenWithUsrBuilder.Build()

	loginResponse := models.PlayerLoginResponseModel{
		Profile: profile,
		Token: token,
	}

	return loginResponse, nil
}
