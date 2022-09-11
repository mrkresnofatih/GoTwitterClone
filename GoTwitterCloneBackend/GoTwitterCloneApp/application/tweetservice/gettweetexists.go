package tweetservice

import (
	"context"
	"errors"
	"log"
	"mrkresnofatihdev/apps/gotwittercloneapp/application"
	"mrkresnofatihdev/apps/gotwittercloneapp/models"
)

func GetTweetExists(ctx context.Context, getRequest models.TweetGetRequestModel) (bool, error) {
	fireStr := application.GetFirestoreInstance()

	foundTweet, err := fireStr.
		Collection(TweetsCollectionKeyFormat).
		Doc(getRequest.TweetId).
		Get(ctx)
	if err != nil {
		log.Println("get tweet failed: " + err.Error())
		return false, errors.New("get tweet results in error")
	}
	return foundTweet.Exists(), nil
}
