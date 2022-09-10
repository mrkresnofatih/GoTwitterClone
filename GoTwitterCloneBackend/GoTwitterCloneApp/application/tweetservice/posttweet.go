package tweetservice

import (
	"context"
	"errors"
	"fmt"
	"github.com/google/uuid"
	"log"
	"mrkresnofatihdev/apps/gotwittercloneapp/application"
	"mrkresnofatihdev/apps/gotwittercloneapp/application/playerservice"
	"mrkresnofatihdev/apps/gotwittercloneapp/models"
	"time"
)

const tweetsCollectionKeyFormat = "tweets#%s"
const tweetKeyFormat = "tweet#%s#%s"

func PostTweet(ctx context.Context, createRequest models.TweetCreateRequestModel, username string) (models.Tweet, error) {
	fireStr := application.GetFirestoreInstance()

	userTweetsCollectionKey := fmt.Sprintf(tweetsCollectionKeyFormat, username)
	tweetKey := fmt.Sprintf(tweetKeyFormat, username, uuid.New().String())

	profile, err := playerservice.GetPlayerMinimumProfile(ctx, username)
	if err != nil {
		log.Println("get profile failed!")
		return *new(models.Tweet), errors.New("failed_create_tweet_get_profile_failed")
	}

	newTweet := models.Tweet{
		TweetId:      tweetKey,
		Username:     username,
		TweetType:    models.SimpleTweetType,
		Message:      createRequest.Message,
		ImageURL:     createRequest.ImageURL,
		CreatedAt:    time.Now(),
		ParentTweet:  nil,
		AvatarURL:    profile.AvatarURL,
		QuoteCount:   0,
		ReplyCount:   0,
		RetweetCount: 0,
	}

	_, err = fireStr.
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
