package followservice

import (
	"context"
	"fmt"
	"mrkresnofatihdev/apps/gotwittercloneapp/application"
	"mrkresnofatihdev/apps/gotwittercloneapp/models"
)

func GetFollowerListExists(ctx context.Context, model models.FollowRequestModel) bool {
	fireStr := application.GetFirestoreInstance()
	followerListKey := fmt.Sprintf(followerListKeyFormat, model.Username, model.FollowerUsername[:6])

	followerList, err := fireStr.
		Collection(followerListCollectionName).
		Doc(followerListKey).
		Get(ctx)
	if err != nil {
		return false
	}
	return followerList.Exists()
}

func GetFollowingListExists(ctx context.Context, model models.FollowRequestModel) bool {
	fireStr := application.GetFirestoreInstance()
	followingListKey := fmt.Sprintf(followingListKeyFormat, model.FollowerUsername, model.Username[:6])

	followingList, err := fireStr.
		Collection(followingListCollectionName).
		Doc(followingListKey).
		Get(ctx)
	if err != nil {
		return false
	}
	return followingList.Exists()
}
