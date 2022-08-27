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
	log.Println("firestore initialized")
	followerListKey := fmt.Sprintf(followerListKeyFormat, model.Username)
	startsWithKey := fmt.Sprintf(startsWithKeyFormat, model.FollowerUsername[:1]) // increase for more follower upper limit

	followerList := models.FollowerList{
		Username: model.Username,
		FollowerList: map[string]bool{
			model.FollowerUsername : true,
		},
		StartsWith: model.FollowerUsername[:1],
	}

	docRef := fireStr.
		Collection(followerListKey).
		Doc(startsWithKey)
	_, err := docRef.Set(ctx, followerList)
	if err != nil {
		log.Println("something went wrong")
		return errors.New("failed to create new follower list")
	}
	return nil
}

func InitializeFollowingList(ctx context.Context, model models.FollowRequestModel) error {
	fireStr := application.GetFirestoreInstance()
	followingListKey := fmt.Sprintf(followingListKeyFormat, model.FollowerUsername)
	startsWithKey := fmt.Sprintf(startsWithKeyFormat, model.Username[:1])

	followingList := models.FollowingList{
		Username: model.FollowerUsername,
		FollowingList: map[string]bool{
			model.Username : true,
		},
		StartsWith: model.Username[:1],
	}

	_, err := fireStr.
		Collection(followingListKey).
		Doc(startsWithKey).
		Set(ctx, followingList)
	if err != nil {
		return errors.New("failed to create new following list")
	}
	return nil
}