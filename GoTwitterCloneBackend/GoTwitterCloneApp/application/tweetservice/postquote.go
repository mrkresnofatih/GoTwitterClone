package tweetservice

import (
	"cloud.google.com/go/firestore"
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

func PostQuote(ctx context.Context, quoteReq models.TweetReplyRequestModel, username string) (models.Tweet, error) {
	returnDefaultsOnError := func(err error) (models.Tweet, error) {
		return *new(models.Tweet), err
	}

	getTweetReq := models.TweetGetRequestModel{
		TweetId: quoteReq.TweetId,
	}
	targetTweetExists, err := GetTweetExists(ctx, getTweetReq)
	if err != nil {
		return returnDefaultsOnError(err)
	}

	if !targetTweetExists {
		log.Println("target tweet doesnt exist")
		return returnDefaultsOnError(err)
	}

	targetTweet, err := GetTweet(ctx, getTweetReq)
	if err != nil {
		log.Println("failed to get tweet")
		return returnDefaultsOnError(err)
	}

	fireStr := application.GetFirestoreInstance()
	tweetKey := fmt.Sprintf(TweetKeyFormat, username, uuid.New().String())

	var parentTweet *models.Tweet
	switch targetTweet.TweetType {
	case models.ReplyTweetType:
		fallthrough
	case models.QuoteTweetType:
		parentTweet = constructParentForQuoteTweetTypeQuoteReply(targetTweet)
	case models.RetweetTweetType:
		log.Println("cannot quote a retweet")
		return returnDefaultsOnError(errors.New("cannot quote a retweet"))
	default:
		parentTweet = &targetTweet
	}

	profile, err := playerservice.GetPlayerMinimumProfile(ctx, username)
	if err != nil {
		log.Println("get profile failed!")
		return returnDefaultsOnError(errors.New("failed_create_tweet_get_profile_failed"))
	}

	parentTweet.QuoteCount++
	newTweet := models.Tweet{
		TweetId:      tweetKey,
		Username:     username,
		TweetType:    models.QuoteTweetType,
		Message:      quoteReq.Message,
		ImageURL:     quoteReq.ImageURL,
		CreatedAt:    fmt.Sprintf("%015d", time.Now().UnixMilli()),
		ParentTweet:  parentTweet,
		AvatarURL:    profile.AvatarURL,
		ReplyCount:   0,
		QuoteCount:   0,
		RetweetCount: 0,
	}

	log.Println("attempt saving quote to db")
	_, err = fireStr.
		Collection(TweetsCollectionKeyFormat).
		Doc(tweetKey).
		Set(ctx, newTweet)
	if err != nil {
		log.Println("something went wrong while saving tweet")
		return returnDefaultsOnError(errors.New("failed to save quote tweet"))
	}
	log.Println("finished saving quote to db")
	return newTweet, nil
}

func constructParentForQuoteTweetTypeQuoteReply(targetTweet models.Tweet) *models.Tweet {
	currentParentTweet := &targetTweet
	if currentParentTweet.ParentTweet != nil {
		currentParentTweet.ParentTweet.ParentTweet = nil
	}
	return currentParentTweet
}

func IncrementQuoteCountOfTargetTweet(ctx context.Context, targetTweet models.TweetReplyRequestModel) error {
	log.Println("attempting to increment quote count of target tweet")

	fireStr := application.GetFirestoreInstance()
	_, err := fireStr.
		Collection(TweetsCollectionKeyFormat).
		Doc(targetTweet.TweetId).
		Update(ctx, []firestore.Update{
			{
				Path:  fmt.Sprintf("quoteCount"),
				Value: firestore.Increment(1),
			},
		})
	if err != nil {
		log.Println("failed to increment quote count")
		return errors.New("increment quote count failed")
	}

	log.Println("increment success")

	return nil
}
