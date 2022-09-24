package followservice

import (
	"context"
	"fmt"
	"log"
	"mrkresnofatihdev/apps/gotwittercloneapp/application"
	"mrkresnofatihdev/apps/gotwittercloneapp/models"
)

func GetIsAlreadyFollowing(ctx context.Context, req models.FollowRequestModel) bool {
	followerListExists := GetFollowerListExists(ctx, req)
	if !followerListExists {
		return false
	}

	fireStr := application.GetFirestoreInstance()

	followerListKey := fmt.Sprintf(followerListKeyFormat, req.Username, req.FollowerUsername[:1])

	followerListData, err := fireStr.
		Collection(followerListCollectionName).
		Doc(followerListKey).
		Get(ctx)
	if err != nil {
		log.Println("something went wrong while fetching follower list data, assumed already following")
		return true
	}

	var followerList models.FollowerList
	err = followerListData.DataTo(&followerList)
	if err != nil {
		log.Println("something went wrong while parsing firebase doc to model")
		return true
	}

	if _, ok := followerList.FollowerList[req.FollowerUsername]; ok {
		log.Println("follower username key found! already following is true!")
		return true
	}
	return false
}
