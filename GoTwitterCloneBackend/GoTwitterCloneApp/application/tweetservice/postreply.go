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

func PostReply(ctx context.Context, replyReq models.TweetReplyRequestModel, username string) (models.Tweet, error) {
	returnDefaultsOnError := func(err error) (models.Tweet, error) {
		return *new(models.Tweet), err
	}

	getTweetReq := models.TweetGetRequestModel{
		TweetId: replyReq.TweetId,
	}
	targetTweetExists, err := GetTweetExists(ctx, getTweetReq)
	if err != nil {
		return returnDefaultsOnError(err)
	}

	if !targetTweetExists {
		log.Println("target tweet doesnt exist! will return error")
		return returnDefaultsOnError(errors.New("target tweet doesnt exist"))
	}

	targetTweet, err := GetTweet(ctx, getTweetReq)
	if err != nil {
		log.Println("failed to get tweet")
		return returnDefaultsOnError(errors.New("failed to get target tweet for reply"))
	}

	fireStr := application.GetFirestoreInstance()
	tweetKey := fmt.Sprintf(TweetKeyFormat, username, uuid.New().String())

	var parentTweet *models.Tweet
	switch targetTweet.TweetType {
	case models.QuoteTweetType:
		parentTweet = constructParentForReplyTweetTypeQuote(targetTweet)
	case models.RetweetTweetType:
		log.Println("cannot reply a retweet type! will continue to fail!")
		return returnDefaultsOnError(errors.New("cannot reply a retweet"))
	case models.ReplyTweetType:
		parentTweet = constructParentForReplyTweetTypeReply(targetTweet)
	default:
		parentTweet = &targetTweet
	}

	profile, err := playerservice.GetPlayerMinimumProfile(ctx, username)
	if err != nil {
		log.Println("get profile failed!")
		return returnDefaultsOnError(errors.New("failed_create_tweet_get_profile_failed"))
	}

	parentTweet.ReplyCount++
	newTweet := models.Tweet{
		TweetId:      tweetKey,
		Username:     username,
		TweetType:    models.ReplyTweetType,
		Message:      replyReq.Message,
		ImageURL:     replyReq.ImageURL,
		CreatedAt:    fmt.Sprintf("%015d", time.Now().UnixMilli()),
		ParentTweet:  parentTweet,
		AvatarURL:    profile.AvatarURL,
		ReplyCount:   0,
		RetweetCount: 0,
		QuoteCount:   0,
	}

	log.Println("attempt saving reply to db")
	_, err = fireStr.
		Collection(TweetsCollectionKeyFormat).
		Doc(tweetKey).
		Set(ctx, newTweet)
	if err != nil {
		log.Println("something went wrong while saving tweet")
		return *new(models.Tweet), errors.New("failed_create_tweet: " + err.Error())
	}
	log.Println("finished saving reply to db")
	return newTweet, nil
}

func constructParentForReplyTweetTypeReply(targetTweet models.Tweet) *models.Tweet {
	maxTraversal := 4
	traversalCount := 0
	currentParentTweet := &targetTweet
	log.Println("starting reply parent-tweet-with-reply-type")
	for traversalCount <= maxTraversal {
		if currentParentTweet == nil {
			log.Println("currentParentTweet is nil! will continue to break from loop")
			break
		}

		if currentParentTweet.TweetType != models.ReplyTweetType {
			log.Println("currentParentTweet tweetType != replyType, will continue to break from loop")
			break
		}

		if traversalCount+1 == maxTraversal {
			log.Println("traversal count reached it's limit, parent will be nil")
			currentParentTweet.ParentTweet = nil
		} else {
			log.Println("traverse to parent")
			currentParentTweet = currentParentTweet.ParentTweet
		}
		traversalCount++
	}

	return &targetTweet
}

func constructParentForReplyTweetTypeQuote(targetTweet models.Tweet) *models.Tweet {
	currentParentTweet := &targetTweet
	if currentParentTweet.ParentTweet != nil {
		currentParentTweet.ParentTweet.ParentTweet = nil
	}
	return currentParentTweet
}

func IncrementReplyCountOfTargetTweet(ctx context.Context, targetTweet models.TweetReplyRequestModel) error {
	log.Println("attempting to increment reply count of target tweet")

	fireStr := application.GetFirestoreInstance()
	_, err := fireStr.
		Collection(TweetsCollectionKeyFormat).
		Doc(targetTweet.TweetId).
		Update(ctx, []firestore.Update{
			{
				Path:  fmt.Sprintf("replyCount"),
				Value: firestore.Increment(1),
			},
		})
	if err != nil {
		log.Println("failed to increment reply count")
		return errors.New("increment reply count failed")
	}

	log.Println("increment success")

	return nil
}
