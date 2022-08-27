package playerservice

import (
	"context"
	"errors"
	"fmt"
	"mrkresnofatihdev/apps/gotwittercloneapp/application"
	"mrkresnofatihdev/apps/gotwittercloneapp/models"
)

func GetPlayerCredentials(ctx context.Context, username string) (models.PlayerCredentials, error) {
	returnDefaultsOnError := func(err error) (models.PlayerCredentials, error) {
		return *new(models.PlayerCredentials), err
	}

	fireStr := application.GetFirestoreInstance()

	playerKey := fmt.Sprintf(PlayerKeyFormat, username)
	playerCredentialsKey := fmt.Sprintf(playerCredentialsKeyFormat, username)

	foundRecord, err := fireStr.
		Collection(playerKey).
		Doc(playerCredentialsKey).
		Get(ctx)
	if err != nil {
		return returnDefaultsOnError(errors.New("player_credentials_not_found"))
	}

	var playerCredential models.PlayerCredentials
	err = foundRecord.DataTo(&playerCredential)
	if err != nil {
		return returnDefaultsOnError(errors.New("found_player_credentials_parse_failed"))
	}
	return playerCredential, nil
}
