package playerservice

import (
	"context"
	"fmt"
	"mrkresnofatihdev/apps/gotwittercloneapp/application"
)

func GetPlayerExists(ctx context.Context, username string) bool {
	fireStr := application.GetFirestoreInstance()
	playerKey := fmt.Sprintf(playerKeyFormat, username)
	playerProfileKey := fmt.Sprintf(playerProfileKeyFormat, username)
	player, err := fireStr.
		Collection(playerKey).
		Doc(playerProfileKey).
		Get(ctx)
	if err != nil {
		return false
	}
	return player.Exists()
}
