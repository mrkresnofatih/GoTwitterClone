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

const TweetsCollectionName = "tweets"
const TweetKeyFormat = "tweet#%s#%s"

func PostTweet(ctx context.Context, createRequest models.TweetCreateRequestModel, username string) (models.Tweet, error) {
	fireStr := application.GetFirestoreInstance()

	tweetKey := fmt.Sprintf(TweetKeyFormat, username, uuid.New().String())

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
		CreatedAt:    fmt.Sprintf("%015d", time.Now().UnixMilli()),
		ParentTweet:  nil,
		AvatarURL:    profile.AvatarURL,
		QuoteCount:   0,
		ReplyCount:   0,
		RetweetCount: 0,
	}

	_, err = fireStr.
		Collection(TweetsCollectionName).
		Doc(tweetKey).
		Set(ctx, newTweet)
	if err != nil {
		log.Println("something went wrong while saving tweet")
		return *new(models.Tweet), errors.New("failed_create_tweet: " + err.Error())
	}
	log.Println("successfully saved tweet!")
	return newTweet, nil
}
