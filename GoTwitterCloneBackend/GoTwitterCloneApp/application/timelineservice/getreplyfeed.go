package timelineservice

import (
	"cloud.google.com/go/firestore"
	"context"
	"fmt"
	"google.golang.org/api/iterator"
	"log"
	"mrkresnofatihdev/apps/gotwittercloneapp/application"
	"mrkresnofatihdev/apps/gotwittercloneapp/application/tweetservice"
	"mrkresnofatihdev/apps/gotwittercloneapp/models"
)

func GetReplyFeed(ctx context.Context, query models.ReplyFeedQueryModel) ([]models.Tweet, error) {
	returnDefaultOnError := func(err error) ([]models.Tweet, error) {
		return *new([]models.Tweet), err
	}

	fireStr := application.GetFirestoreInstance()

	var tweets []models.Tweet

	log.Println("get iter")

	startAt := fmt.Sprintf("%015d", query.StartAt)
	maxLimit := 20
	if query.Limit > maxLimit {
		log.Println("query-limit reply-feed too high! will default to max-limit")
		query.Limit = maxLimit
	}
	iter := fireStr.
		Collection(tweetservice.TweetReplyRecordCollectionName).
		Where("tweetId", "==", query.TweetId).
		Where("createdAt", ">", startAt).
		OrderBy("createdAt", firestore.Desc).
		Limit(query.Limit).
		Documents(ctx)
	for {
		doc, err := iter.Next()
		if err == iterator.Done {
			log.Println("iter done")
			break
		}

		if err != nil {
			log.Println("error at iter")
			return returnDefaultOnError(err)
		}

		var replyRecord models.TweetReplyRecordModel
		err = doc.DataTo(&replyRecord)
		if err != nil {
			log.Println("error cannot parse doc to reply-record")
			continue
		}

		tweetGetReq := models.TweetGetRequestModel{TweetId: replyRecord.ReplyTweetId}
		tweetExists, err := tweetservice.GetTweetExists(ctx, tweetGetReq)
		if err != nil {
			log.Println("something went wrong when get-tweet-exists!")
			continue
		}

		if !tweetExists {
			log.Println("tweet doesn't exist")
			continue
		}

		tweet, err := tweetservice.GetTweet(ctx, tweetGetReq)
		if err != nil {
			log.Println("error while getting tweet")
			continue
		}

		tweets = append(tweets, tweet)
	}

	return tweets, nil
}
