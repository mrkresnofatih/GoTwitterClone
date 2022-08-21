package playerservice

import (
	"context"
	"errors"
	"fmt"
	"mrkresnofatihdev/apps/gotwittercloneapp/application"
	"mrkresnofatihdev/apps/gotwittercloneapp/models"
)

func GetPlayerMinimumProfile(ctx context.Context, username string) (models.PlayerGetMinimumProfileResponseModel, error) {
	returnDefaultsOnError := func(err error) (models.PlayerGetMinimumProfileResponseModel, error) {
		return *new(models.PlayerGetMinimumProfileResponseModel), err
	}

	fireStr := application.GetFirestoreInstance()

	playerKey := fmt.Sprintf(playerKeyFormat, username)
	foundRecord, err := fireStr.
		Collection(playerKey).
		Doc(playerKey).
		Get(ctx)
	if err != nil || !foundRecord.Exists() {
		return returnDefaultsOnError(errors.New("player_not_found"))
	}

	var playerInfo models.Player
	err = foundRecord.DataTo(&playerInfo)
	if err != nil {
		return returnDefaultsOnError(errors.New("found_player_parse_failed"))
	}
	return playerInfo.ToGetMinimumProfileResponseModel(), nil
}
