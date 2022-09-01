package tweetservice

import (
	"context"
	"errors"
	"fmt"
	"log"
	"mrkresnofatihdev/apps/gotwittercloneapp/application"
	"mrkresnofatihdev/apps/gotwittercloneapp/models"
)

func GetTweetExists(ctx context.Context, getRequest models.TweetGetRequestModel) (bool, error) {
	fireStr := application.GetFirestoreInstance()

	userTweetsCollectionKey := fmt.Sprintf(tweetsCollectionKeyFormat, getRequest.Username)

	foundTweet, err := fireStr.
		Collection(userTweetsCollectionKey).
		Doc(getRequest.TweetId).
		Get(ctx)
	if err != nil {
		log.Println("get tweet failed: " + err.Error())
		return false, errors.New("get tweet results in error")
	}
	return foundTweet.Exists(), nil
}
