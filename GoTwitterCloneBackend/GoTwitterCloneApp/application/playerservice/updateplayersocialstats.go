package playerservice

import (
	"cloud.google.com/go/firestore"
	"context"
	"errors"
	"fmt"
	"mrkresnofatihdev/apps/gotwittercloneapp/application"
	"mrkresnofatihdev/apps/gotwittercloneapp/models"
)

func UpdatePlayerSocialStats(ctx context.Context, updateRequest models.PlayerUpdateSocialStatsRequestModel) error {
	playerSocialStatsKey := fmt.Sprintf(PlayerSocialStatsKeyFormat, updateRequest.Username)
	playerKey := fmt.Sprintf(PlayerKeyFormat, updateRequest.Username)

	fireStr := application.GetFirestoreInstance()

	targetDoc := fireStr.Collection(playerKey).Doc(playerSocialStatsKey)
	var updateObject []firestore.Update

	switch updateRequest.UpdateType {
	case models.IncrementFollowerUpdateSocialStatsType:
		updateObject = append(updateObject, firestore.Update{
			Path: "numOfFollowers",
			Value: firestore.Increment(1),
		})
	case models.DecrementFollowerUpdateSocialStatsType:
		updateObject = append(updateObject, firestore.Update{
			Path: "numOfFollowers",
			Value: firestore.Increment(-1),
		})
	case models.IncrementFollowingUpdateSocialStatsType:
		updateObject = append(updateObject, firestore.Update{
			Path: "numOfFollowings",
			Value: firestore.Increment(1),
		})
	case models.DecrementFollowingUpdateSocialStatsType:
		updateObject = append(updateObject, firestore.Update{
			Path: "numOfFollowings",
			Value: firestore.Increment(-1),
		})
	}

	_, err := targetDoc.Update(ctx, updateObject)
	if err != nil {
		return errors.New("failed to update player social stats")
	}
	return nil
}