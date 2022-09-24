package eventhandlers

import (
	"context"
	"log"
	"mrkresnofatihdev/apps/gotwittercloneapp/application/timelineservice"
)

func RunGetHomeQueryEventHandler(ctx context.Context, username string) {
	err := timelineservice.RunHomeFeedQuery(ctx, username)
	if err != nil {
		log.Println("run-home-query failed due to error: " + err.Error())
	}
}
