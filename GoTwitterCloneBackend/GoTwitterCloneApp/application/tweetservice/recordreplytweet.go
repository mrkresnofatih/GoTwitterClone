package tweetservice

import (
	"context"
	"fmt"
	"log"
	"mrkresnofatihdev/apps/gotwittercloneapp/application"
	"mrkresnofatihdev/apps/gotwittercloneapp/models"
	"strings"
	"time"
)

const TweetReplyRecordCollectionName = "replyRecords"

func RecordReplyTweet(ctx context.Context, replyTweet models.Tweet) {
	replyRecordExists, err := getTweetReplyRecordExists(ctx, replyTweet.TweetId)
	if err != nil {
		log.Println("get-tweet-reply-record-exists failed")
		return
	}

	if replyRecordExists {
		log.Println("reply-record already exists")
		return
	}

	fireStr := application.GetFirestoreInstance()

	newReplyRecord := models.TweetReplyRecordModel{
		TweetId:      replyTweet.ParentTweet.TweetId,
		ReplyTweetId: replyTweet.TweetId,
		CreatedAt:    fmt.Sprintf("%015d", time.Now().UnixMilli()),
	}

	_, err = fireStr.
		Collection(TweetReplyRecordCollectionName).
		Doc(replyTweet.TweetId).
		Set(ctx, newReplyRecord)
	if err != nil {
		log.Println("something went wrong while saving reply-record")
		return
	}
	log.Println("successfully saved reply-record!")
	return
}

func getTweetReplyRecordExists(ctx context.Context, replyTweetId string) (bool, error) {
	fireStr := application.GetFirestoreInstance()
	record, err := fireStr.
		Collection(TweetReplyRecordCollectionName).
		Doc(replyTweetId).
		Get(ctx)
	if err != nil {
		if strings.HasPrefix(err.Error(), "rpc error: code = NotFound") {
			log.Println("tweet reply record not found")
			return false, nil
		}
		return false, err
	}

	return record.Exists(), nil
}
