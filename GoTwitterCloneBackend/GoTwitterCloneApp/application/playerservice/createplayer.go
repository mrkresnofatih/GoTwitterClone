package playerservice

import (
	"context"
	"errors"
	"fmt"
	"mrkresnofatihdev/apps/gotwittercloneapp/application"
	"mrkresnofatihdev/apps/gotwittercloneapp/models"
	"mrkresnofatihdev/apps/gotwittercloneapp/utils"
	"time"
)

const defaultAvatarUrl = "https://cdn.iconscout.com/icon/free/png-256/avatar-372-456324.png"
const defaultBioFormat = "Hi, I'm %s!"

const playerKeyFormat = "players#%s"
const playerProfileKeyFormat = "players-profile#%s"
const playerCredentialsKeyFormat = "players-cred#%s"

func CreatePlayer(ctx context.Context, createRequest models.PlayerCreateRequestModel) (models.PlayerCreateResponseModel, error) {
	returnDefaultsOnError := func(err error) (models.PlayerCreateResponseModel, error) {
		return *new(models.PlayerCreateResponseModel), err
	}

	isPlayerUsernameTaken := GetPlayerExists(ctx, createRequest.Username)
	if isPlayerUsernameTaken {
		return returnDefaultsOnError(errors.New("player_username_taken"))
	}

	fireStr := application.GetFirestoreInstance()
	playerKey := fmt.Sprintf(playerKeyFormat, createRequest.Username)
	playerProfileKey := fmt.Sprintf(playerProfileKeyFormat, createRequest.Username)
	playerCredKey := fmt.Sprintf(playerCredentialsKeyFormat, createRequest.Username)

	newPlayerInfo := models.Player{
		Username: createRequest.Username,
		FullName: createRequest.FullName,
		AvatarURL: defaultAvatarUrl,
		CreatedAt: time.Now().UnixMilli(),
		Bio: fmt.Sprintf(defaultBioFormat, createRequest.Username),
		Email: createRequest.Email,
	}

	hashUtil := utils.HashUtility{
		HashStrategy: &utils.HashUserPasswordStrategy{},
	}

	newPlayerCredentials := models.PlayerCredentials{
		Username: createRequest.Username,
		Email: createRequest.Email,
		Password: hashUtil.GetHashData(createRequest.Password),
	}

	_, err := fireStr.
		Collection(playerKey).
		Doc(playerCredKey).
		Set(ctx, newPlayerCredentials)
	if err != nil {
		return returnDefaultsOnError(err)
	}

	_, err = fireStr.
		Collection(playerKey).
		Doc(playerProfileKey).
		Set(ctx, newPlayerInfo)
	if err != nil {
		return returnDefaultsOnError(err)
	}
	return newPlayerInfo.ToCreateResponseModel(), nil
}
