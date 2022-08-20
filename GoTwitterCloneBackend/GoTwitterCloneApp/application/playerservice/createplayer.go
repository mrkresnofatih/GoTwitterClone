package playerservice

import (
	"context"
	"fmt"
	"mrkresnofatihdev/apps/gotwittercloneapp/application"
	"mrkresnofatihdev/apps/gotwittercloneapp/models"
	"mrkresnofatihdev/apps/gotwittercloneapp/utils"
	"time"
)

const defaultAvatarUrl = "https://cdn.iconscout.com/icon/free/png-256/avatar-372-456324.png"
const defaultBioFormat = "Hi, I'm %s!"
const playerCollectionFormat = "players#%s"
const playerCredentialsCollectionFormat = "players-cred#%s"

func CreatePlayer(ctx context.Context, createRequest models.PlayerCreateRequestModel) (models.PlayerCreateResponseModel, error) {
	returnDefaultsOnError := func(err error) (models.PlayerCreateResponseModel, error) {
		return *new(models.PlayerCreateResponseModel), err
	}

	fireStr := application.GetFirestoreInstance()

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

	playerCredCollectionName := fmt.Sprintf(playerCredentialsCollectionFormat, createRequest.Username)
	_, err := fireStr.
		Collection(playerCredCollectionName).
		Doc(newPlayerInfo.Username).
		Set(ctx, newPlayerCredentials)
	if err != nil {
		return returnDefaultsOnError(err)
	}

	playerCollectionName := fmt.Sprintf(playerCollectionFormat, createRequest.Username)
	_, err = fireStr.
		Collection(playerCollectionName).
		Doc(newPlayerInfo.Username).
		Set(ctx, newPlayerInfo)
	if err != nil {
		return returnDefaultsOnError(err)
	}
	return newPlayerInfo.ToCreateResponseModel(), nil
}
