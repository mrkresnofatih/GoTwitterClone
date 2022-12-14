package tweetservice

import (
	"context"
	"log"
	"mrkresnofatihdev/apps/gotwittercloneapp/application"
	"mrkresnofatihdev/apps/gotwittercloneapp/models"
)

func GetTweet(ctx context.Context, getRequest models.TweetGetRequestModel) (models.Tweet, error) {
	returnDefaultsOnError := func(err error) (models.Tweet, error) {
		return *new(models.Tweet), err
	}

	fireStr := application.GetFirestoreInstance()

	foundTweetData, err := fireStr.
		Collection(TweetsCollectionName).
		Doc(getRequest.TweetId).
		Get(ctx)
	if err != nil {
		log.Println("get tweet failed: " + err.Error())
		return returnDefaultsOnError(err)
	}

	var foundTweet models.Tweet
	err = foundTweetData.DataTo(&foundTweet)
	if err != nil {
		log.Println("parse tweet to model failed")
		return returnDefaultsOnError(err)
	}
	return foundTweet, nil
}
