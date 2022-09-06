package eventhandlers

import (
	"context"
	"log"
	"mrkresnofatihdev/apps/gotwittercloneapp/application/followservice"
	"mrkresnofatihdev/apps/gotwittercloneapp/models"
)

func FollowEventHandler(followRequest models.FollowRequestModel) {
	ctx := context.Background()
	err := followservice.StartFollowing(ctx, followRequest)
	if err != nil {
		log.Println("start following failed!")
	}
}

func UnfollowEventHandler(followRequest models.FollowRequestModel) {
	ctx := context.Background()
	err := followservice.StopFollowing(ctx, followRequest)
	if err != nil {
		log.Println("stop following failed")
	}
}

const FollowEventHandlerName = "FollowEvtHandler"
const UnfollowEventHandlerName = "UnfollowEvtHandler"
