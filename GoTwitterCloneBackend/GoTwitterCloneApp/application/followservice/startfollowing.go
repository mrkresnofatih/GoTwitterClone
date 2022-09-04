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

const followerListKeyFormat = "follower-list#%s"
const followingListKeyFormat = "following-list#%s"
const startsWithKeyFormat = "startsWith#%s"

func StartFollowing(ctx context.Context, followRequest models.FollowRequestModel) error {
	alreadyFollowing := GetIsAlreadyFollowing(ctx, followRequest)
	if alreadyFollowing {
		log.Println("already following!")
		return nil
	}

	followerListExists := GetFollowerListExists(ctx, followRequest)
	if !followerListExists {
		log.Println("follower list doesn't exist")
		err := InitializeFollowerList(ctx, followRequest)
		if err != nil {
			log.Println("InitializeFollowerList failed: " + err.Error())
			return err
		}
	} else {
		err := populateFollowerList(ctx, followRequest)
		if err != nil {
			log.Println("populateFollowersList failed: " + err.Error())
			return err
		}
	}

	userIncrementNumOfFollowersRequest := models.PlayerUpdateSocialStatsRequestModel{
		Username: followRequest.Username,
		UpdateType: models.IncrementFollowerUpdateSocialStatsType,
	}
	err := playerservice.UpdatePlayerSocialStats(ctx, userIncrementNumOfFollowersRequest)
	if err != nil {
		log.Println("UpdatePlayerSocialStats failed: " + err.Error())
		return err
	}

	followingListExists := GetFollowingListExists(ctx, followRequest)
	if !followingListExists {
		err := InitializeFollowingList(ctx, followRequest)
		if err != nil {
			log.Println("InitializeFollowingList failed: " + err.Error())
			return err
		}
	} else {
		err := populateFollowingList(ctx, followRequest)
		if err != nil {
			log.Println("populateFollowingList failed: " + err.Error())
			return err
		}
	}

	followerIncrementNumOfFollowingsRequest := models.PlayerUpdateSocialStatsRequestModel{
		Username: followRequest.FollowerUsername,
		UpdateType: models.IncrementFollowingUpdateSocialStatsType,
	}
	err = playerservice.UpdatePlayerSocialStats(ctx, followerIncrementNumOfFollowingsRequest)
	if err != nil {
		log.Println("UpdatePlayerSocialStats failed: " + err.Error())
		return err
	}

	return nil
}

func populateFollowerList(ctx context.Context, followRequest models.FollowRequestModel) error {
	followerListKey := fmt.Sprintf(followerListKeyFormat, followRequest.Username)
	startsWithKey := fmt.Sprintf(startsWithKeyFormat, followRequest.FollowerUsername[:1])

	fireStr := application.GetFirestoreInstance()

	_, err := fireStr.Collection(followerListKey).Doc(startsWithKey).Update(ctx, []firestore.Update{
		{
			Path: fmt.Sprintf("followerList.%s", followRequest.FollowerUsername),
			Value: true,
		},
	})
	if err != nil {
		log.Println("Add follower to follower-list failed: " + err.Error())
		return errors.New("populate follower to target follower-list failed")
	}
	return nil
}

func populateFollowingList(ctx context.Context, followRequest models.FollowRequestModel) error {
	followingListKey := fmt.Sprintf(followingListKeyFormat, followRequest.FollowerUsername)
	startsWithKey := fmt.Sprintf(startsWithKeyFormat, followRequest.Username[:1])

	fireStr := application.GetFirestoreInstance()

	_, err := fireStr.Collection(followingListKey).Doc(startsWithKey).Update(ctx, []firestore.Update{
		{
			Path: fmt.Sprintf("followingList.%s", followRequest.Username),
			Value: true,
		},
	})
	if err != nil {
		return errors.New("populate following to target following-list failed")
	}
	return nil
}