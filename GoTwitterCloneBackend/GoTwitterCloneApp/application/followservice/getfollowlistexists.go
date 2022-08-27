package followservice

import (
	"context"
	"fmt"
	"mrkresnofatihdev/apps/gotwittercloneapp/application"
	"mrkresnofatihdev/apps/gotwittercloneapp/models"
)

func GetFollowerListExists(ctx context.Context, model models.FollowRequestModel) bool {
	fireStr := application.GetFirestoreInstance()
	followerListKey := fmt.Sprintf(followerListKeyFormat, model.Username)
	startsWithKey := fmt.Sprintf(startsWithKeyFormat, model.FollowerUsername[:1]) // increase for more follower upper limit

	followerList, err := fireStr.
		Collection(followerListKey).
		Doc(startsWithKey).
		Get(ctx)
	if err != nil {
		return false
	}
	return followerList.Exists()
}

func GetFollowingListExists(ctx context.Context, model models.FollowRequestModel) bool {
	fireStr := application.GetFirestoreInstance()
	followingListKey := fmt.Sprintf(followingListKeyFormat, model.FollowerUsername)
	startsWithKey := fmt.Sprintf(startsWithKeyFormat, model.Username[:1])

	followingList, err := fireStr.
		Collection(followingListKey).
		Doc(startsWithKey).
		Get(ctx)
	if err != nil {
		return false
	}
	return followingList.Exists()
}