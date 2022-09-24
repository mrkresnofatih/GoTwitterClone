package followservice

import (
	"context"
	"errors"
	"fmt"
	"log"
	"mrkresnofatihdev/apps/gotwittercloneapp/application"
	"mrkresnofatihdev/apps/gotwittercloneapp/models"
)

func InitializeFollowerList(ctx context.Context, model models.FollowRequestModel) error {
	fireStr := application.GetFirestoreInstance()
	followerListKey := fmt.Sprintf(followerListKeyFormat, model.Username, model.FollowerUsername[:6])

	followerList := models.FollowerList{
		Username: model.Username,
		FollowerList: map[string]bool{
			model.FollowerUsername: true,
		},
		StartsWith: model.FollowerUsername[:6],
	}

	docRef := fireStr.
		Collection(followerListCollectionName).
		Doc(followerListKey)
	_, err := docRef.Set(ctx, followerList)
	if err != nil {
		log.Println("something went wrong")
		return errors.New("failed to create new follower list")
	}
	return nil
}

func InitializeFollowingList(ctx context.Context, model models.FollowRequestModel) error {
	fireStr := application.GetFirestoreInstance()
	followingListKey := fmt.Sprintf(followingListKeyFormat, model.FollowerUsername, model.Username[:6])

	followingList := models.FollowingList{
		Username: model.FollowerUsername,
		FollowingList: map[string]bool{
			model.Username: true,
		},
		StartsWith: model.Username[:6],
	}

	_, err := fireStr.
		Collection(followingListCollectionName).
		Doc(followingListKey).
		Set(ctx, followingList)
	if err != nil {
		return errors.New("failed to create new following list")
	}
	return nil
}
