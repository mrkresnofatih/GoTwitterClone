package followservice

import (
	"cloud.google.com/go/firestore"
	"context"
	"errors"
	"fmt"
	"log"
	"mrkresnofatihdev/apps/gotwittercloneapp/application"
	"mrkresnofatihdev/apps/gotwittercloneapp/application/playerservice"
	"mrkresnofatihdev/apps/gotwittercloneapp/models"
)

func StopFollowing(ctx context.Context, followRequest models.FollowRequestModel) error {
	alreadyFollowing := GetIsAlreadyFollowing(ctx, followRequest)
	if !alreadyFollowing {
		log.Println("hasn't followed yet")
		return nil
	}

	returnDefaultsOnError := func(err error) error {
		log.Println("failed: " + err.Error())
		return err
	}

	err := removeFollowerFromFollowerList(ctx, followRequest.Username, followRequest.FollowerUsername)
	if err != nil {
		return returnDefaultsOnError(err)
	}

	err = removeFollowingFromFollowingList(ctx, followRequest.Username, followRequest.FollowerUsername)
	if err != nil {
		return returnDefaultsOnError(err)
	}

	decrementFollowerCountReq := models.PlayerUpdateSocialStatsRequestModel{
		Username: followRequest.Username,
		UpdateType: models.DecrementFollowerUpdateSocialStatsType,
	}
	err = playerservice.UpdatePlayerSocialStats(ctx, decrementFollowerCountReq)
	if err != nil {
		return returnDefaultsOnError(err)
	}

	decrementFollowingCountReq := models.PlayerUpdateSocialStatsRequestModel{
		Username: followRequest.FollowerUsername,
		UpdateType: models.DecrementFollowingUpdateSocialStatsType,
	}
	err = playerservice.UpdatePlayerSocialStats(ctx, decrementFollowingCountReq)
	if err != nil {
		return returnDefaultsOnError(err)
	}

	return nil
}

func removeFollowerFromFollowerList(ctx context.Context, username, followerUsername string) error {
	fireStr := application.GetFirestoreInstance()

	followerListKey := fmt.Sprintf(followerListKeyFormat, username)
	startsWithKey := fmt.Sprintf(startsWithKeyFormat, followerUsername[:1])

	_, err := fireStr.Collection(followerListKey).Doc(startsWithKey).Update(ctx, []firestore.Update{
		{
			Path: fmt.Sprintf("followerList.%s", followerUsername),
			Value: firestore.Delete,
		},
	})
	if err != nil {
		return errors.New("failed to delete follower from follower-list")
	}
	return nil
}

func removeFollowingFromFollowingList(ctx context.Context, username, followerUsername string) error {
	fireStr := application.GetFirestoreInstance()

	followingListKey := fmt.Sprintf(followingListKeyFormat, followerUsername)
	startsWithKey := fmt.Sprintf(startsWithKeyFormat, username[:1])

	_, err := fireStr.
		Collection(followingListKey).
		Doc(startsWithKey).
		Update(ctx, []firestore.Update{
		{
			Path: fmt.Sprintf("followingList.%s", username),
			Value: firestore.Delete,
		},
	})

	if err != nil {
		return errors.New("failed to delete following from following-list")
	}
	return nil
}
