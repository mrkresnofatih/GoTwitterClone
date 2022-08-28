package followservice

import (
	"context"
	"fmt"
	"mrkresnofatihdev/apps/gotwittercloneapp/application"
	"mrkresnofatihdev/apps/gotwittercloneapp/models"
)

func GetIsAlreadyFollowing(ctx context.Context, req models.FollowRequestModel) bool {
	followerListExists := GetFollowerListExists(ctx, req)
	if !followerListExists {
		return false
	}

	fireStr := application.GetFirestoreInstance()

	followerListKey := fmt.Sprintf(followerListKeyFormat, req.Username)
	startsWithKey := fmt.Sprintf(startsWithKeyFormat, req.FollowerUsername[:1])

	followerListData, err := fireStr.
		Collection(followerListKey).
		Doc(startsWithKey).
		Get(ctx)
	if err != nil {
		return true
	}

	var followerList models.FollowerList
	err = followerListData.DataTo(&followerList)
	if err != nil {
		return true
	}

	if _, ok := followerList.FollowerList[req.FollowerUsername]; ok {
		return true
	}
	return false
}
