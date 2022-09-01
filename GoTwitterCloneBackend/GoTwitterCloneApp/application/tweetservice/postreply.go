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

func PostReply(ctx context.Context, replyReq models.TweetReplyRequestModel, username string) (models.Tweet, error) {
	returnDefaultsOnError := func(err error) (models.Tweet, error) {
		return *new(models.Tweet), err
	}

	getTweetReq := models.TweetGetRequestModel{
		TweetId: replyReq.TweetId,
		Username: replyReq.Username,
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

	userTweetsCollectionKey := fmt.Sprintf(tweetsCollectionKeyFormat, username)
	tweetKey := fmt.Sprintf(tweetKeyFormat, username, uuid.New().String())

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

	newTweet := models.Tweet{
		TweetId:     tweetKey,
		Username:    username,
		TweetType:   models.ReplyTweetType,
		Message:     replyReq.Message,
		ImageURL:    replyReq.ImageURL,
		CreatedAt:   time.Now(),
		ParentTweet: parentTweet,
	}

	log.Println("attempt saving reply to db")
	_, err = fireStr.
		Collection(userTweetsCollectionKey).
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

		if traversalCount + 1 == maxTraversal {
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