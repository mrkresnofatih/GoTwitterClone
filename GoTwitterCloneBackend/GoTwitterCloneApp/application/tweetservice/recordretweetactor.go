package tweetservice

import (
	"context"
	"fmt"
	"log"
	"mrkresnofatihdev/apps/gotwittercloneapp/application"
	"mrkresnofatihdev/apps/gotwittercloneapp/models"
)

const retweetKeyFormat = "retweet-actor#%s"
const retweetActorsCollectionName = "retweetActors"

func RecordRetweetActor(ctx context.Context, tweetRetweetActorModel models.TweetRetweetActorModel) error {
	fireStr := application.GetFirestoreInstance()
	retweetKey := fmt.Sprintf(retweetKeyFormat, tweetRetweetActorModel.ActorUsername)

	_, err := fireStr.
		Collection(TweetsCollectionKeyFormat).
		Doc(tweetRetweetActorModel.TweetId).
		Collection(retweetActorsCollectionName).
		Doc(retweetKey).
		Set(ctx, tweetRetweetActorModel)
	if err != nil {
		log.Println("something went wrong while saving retweet actor record")
		return err
	}

	log.Println("successfully saved retweet actor record")
	return nil
}

func RetweetActorRecordExists(ctx context.Context, tweetRetweetActorModel models.TweetRetweetActorModel) bool {
	fireStr := application.GetFirestoreInstance()
	retweetKey := fmt.Sprintf(retweetKeyFormat, tweetRetweetActorModel.ActorUsername)

	tweetRetweetActorRecord, err := fireStr.
		Collection(TweetsCollectionKeyFormat).
		Doc(tweetRetweetActorModel.TweetId).
		Collection(retweetActorsCollectionName).
		Doc(retweetKey).
		Get(ctx)
	if err != nil {
		log.Println("something went wrong while fetching tweet retweet actor record")
		return false
	}
	return tweetRetweetActorRecord.Exists()

}
