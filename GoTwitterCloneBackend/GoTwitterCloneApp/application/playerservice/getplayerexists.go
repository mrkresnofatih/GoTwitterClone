package playerservice

import (
	"context"
	"fmt"
	"mrkresnofatihdev/apps/gotwittercloneapp/application"
)

func GetPlayerExists(ctx context.Context, username string) (bool) {
	fireStr := application.GetFirestoreInstance()
	playerKey := fmt.Sprintf(playerKeyFormat, username)
	player, err := fireStr.
		Collection(playerKey).
		Doc(playerKey).
		Get(ctx)
	if err != nil {
		return false
	}
	return player.Exists()
}
