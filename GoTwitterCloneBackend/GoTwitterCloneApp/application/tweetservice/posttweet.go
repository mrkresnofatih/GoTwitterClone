package tweetservice

import (
	"context"
	"errors"
	"fmt"
	"github.com/google/uuid"
	"log"
	"mrkresnofatihdev/apps/gotwittercloneapp/application"
	"mrkresnofatihdev/apps/gotwittercloneapp/models"
	"time"
)

const tweetsCollectionKeyFormat = "tweets#%s"
const tweetKeyFormat = "tweet#%s#%s"

func PostTweet(ctx context.Context, createRequest models.TweetCreateRequestModel, username string) (models.Tweet, error) {
	fireStr := application.GetFirestoreInstance()

	userTweetsCollectionKey := fmt.Sprintf(tweetsCollectionKeyFormat, username)
	tweetKey := fmt.Sprintf(tweetKeyFormat, username, uuid.New().String())

	newTweet := models.Tweet{
		TweetId:     tweetKey,
		Username:    username,
		TweetType:   models.SimpleTweetType,
		Message:     createRequest.Message,
		ImageURL:    createRequest.ImageURL,
		CreatedAt:   time.Now(),
		ParentTweet: nil,
	}

	_, err := fireStr.
		Collection(userTweetsCollectionKey).
		Doc(tweetKey).
		Set(ctx, newTweet)
	if err != nil {
		log.Println("something went wrong while saving tweet")
		return *new(models.Tweet), errors.New("failed_create_tweet: " + err.Error())
	}
	log.Println("successfully saved tweet!")
	return newTweet, nil
}