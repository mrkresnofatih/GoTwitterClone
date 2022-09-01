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

func PostQuote(ctx context.Context, quoteReq models.TweetReplyRequestModel, username string) (models.Tweet, error) {
	returnDefaultsOnError := func(err error) (models.Tweet, error) {
		return *new(models.Tweet), err
	}

	getTweetReq := models.TweetGetRequestModel{
		TweetId: quoteReq.TweetId,
		Username: quoteReq.Username,
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

	userTweetsCollectionKey := fmt.Sprintf(tweetsCollectionKeyFormat, username)
	tweetKey := fmt.Sprintf(tweetKeyFormat, username, uuid.New().String())

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

	newTweet := models.Tweet{
		TweetId: tweetKey,
		Username: username,
		TweetType: models.QuoteTweetType,
		Message: quoteReq.Message,
		ImageURL: quoteReq.ImageURL,
		CreatedAt: time.Now(),
		ParentTweet: parentTweet,
	}

	log.Println("attempt saving quote to db")
	_, err = fireStr.
		Collection(userTweetsCollectionKey).
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