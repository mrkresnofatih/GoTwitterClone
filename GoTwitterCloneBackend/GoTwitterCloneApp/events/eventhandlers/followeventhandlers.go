package eventhandlers

import (
	"context"
	"log"
	"mrkresnofatihdev/apps/gotwittercloneapp/application/followservice"
	"mrkresnofatihdev/apps/gotwittercloneapp/models"
)

func FollowEventHandler(ctx context.Context, followRequest models.FollowRequestModel) {
	err := followservice.StartFollowing(ctx, followRequest)
	if err != nil {
		log.Println("start following failed!")
	}
}

func UnfollowEventHandler(ctx context.Context, followRequest models.FollowRequestModel) {
	err := followservice.StopFollowing(ctx, followRequest)
	if err != nil {
		log.Println("stop following failed")
	}
}
