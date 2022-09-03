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

func PostRetweet(ctx context.Context, retweetReq models.TweetRetweetRequestModel, username string) (models.Tweet, error) {
	returnDefaultsOnError := func(err error) (models.Tweet, error) {
		return *new(models.Tweet), err
	}

	getTweetReq := models.TweetGetRequestModel{
		TweetId: retweetReq.TweetId,
		Username: retweetReq.Username,
	}
	targetTweetExists, err := GetTweetExists(ctx, getTweetReq)
	if err != nil {
		return returnDefaultsOnError(err)
	}

	if !targetTweetExists {
		log.Println("target tweet for retweet doesnt exist")
		return returnDefaultsOnError(errors.New("target tweet doesnt exist"))
	}

	retweetActorData := models.TweetRetweetActorModel{
		TweetId: retweetReq.TweetId,
		TweetOwnerUsername: retweetReq.Username,
		ActorUsername: username,
	}
	alreadyRetweetedBefore := RetweetActorRecordExists(ctx, retweetActorData)
	if alreadyRetweetedBefore {
		log.Println("already retweeted this before")
		return returnDefaultsOnError(errors.New("already retweeted this before"))
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
	case models.QuoteTweetType:
		fallthrough
	case models.ReplyTweetType:
		parentTweet = constructParentForRetweetTweetTypeReplyQuote(targetTweet)
	case models.RetweetTweetType:
		log.Println("cannot retweet a retweet")
		return returnDefaultsOnError(errors.New("cannot retweet a retweet"))
	default:
		parentTweet = &targetTweet
	}

	newTweet := models.Tweet{
		TweetId:     tweetKey,
		Username:    username,
		TweetType:   models.RetweetTweetType,
		Message:     "",
		ImageURL:    "",
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
		return returnDefaultsOnError(errors.New("failed to save tweet"))
	}
	log.Println("finished saving retweet to db")

	err = RecordRetweetActor(ctx, retweetActorData)
	if err != nil {
		log.Println("failed to save retweet actor data")
		return returnDefaultsOnError(err)
	}
	return newTweet, nil
}

func constructParentForRetweetTweetTypeReplyQuote(targetTweet models.Tweet) *models.Tweet {
	currentParentTweet := &targetTweet
	if currentParentTweet.ParentTweet != nil {
		currentParentTweet.ParentTweet.ParentTweet = nil
	}
	return currentParentTweet
}
