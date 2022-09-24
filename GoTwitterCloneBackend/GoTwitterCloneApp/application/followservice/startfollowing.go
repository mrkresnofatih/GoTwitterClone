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

const followerListKeyFormat = "follower-list#%s#%s"
const followingListKeyFormat = "following-list#%s#%s"

const followerListCollectionName = "follower-list"
const followingListCollectionName = "following-list"

func StartFollowing(ctx context.Context, followRequest models.FollowRequestModel) error {
	targetFollowIsNotSelf := followRequest.Username != followRequest.FollowerUsername
	if !targetFollowIsNotSelf {
		log.Println("target follow is self")
		return errors.New("target follow is self")
	}

	targetFollowExists := playerservice.GetPlayerExists(ctx, followRequest.Username)
	if !targetFollowExists {
		log.Println("target player to follow does not exist!")
		return errors.New("target follow does not exist")
	}

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
		log.Println("initialized follower list successful")
	} else {
		err := populateFollowerList(ctx, followRequest)
		if err != nil {
			log.Println("populateFollowersList failed: " + err.Error())
			return err
		}
	}

	userIncrementNumOfFollowersRequest := models.PlayerUpdateSocialStatsRequestModel{
		Username:   followRequest.Username,
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
		log.Println("initialized following list successful")
	} else {
		err := populateFollowingList(ctx, followRequest)
		if err != nil {
			log.Println("populateFollowingList failed: " + err.Error())
			return err
		}
	}

	followerIncrementNumOfFollowingsRequest := models.PlayerUpdateSocialStatsRequestModel{
		Username:   followRequest.FollowerUsername,
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
	followerListKey := fmt.Sprintf(followerListKeyFormat, followRequest.Username, followRequest.FollowerUsername[:1])

	fireStr := application.GetFirestoreInstance()

	_, err := fireStr.
		Collection(followerListCollectionName).
		Doc(followerListKey).
		Update(ctx, []firestore.Update{
			{
				Path:  fmt.Sprintf("followerList.%s", followRequest.FollowerUsername),
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
	followingListKey := fmt.Sprintf(followingListKeyFormat, followRequest.FollowerUsername, followRequest.Username[:1])

	fireStr := application.GetFirestoreInstance()

	_, err := fireStr.
		Collection(followingListCollectionName).
		Doc(followingListKey).
		Update(ctx, []firestore.Update{
			{
				Path:  fmt.Sprintf("followingList.%s", followRequest.Username),
				Value: true,
			},
		})
	if err != nil {
		return errors.New("populate following to target following-list failed")
	}
	return nil
}
