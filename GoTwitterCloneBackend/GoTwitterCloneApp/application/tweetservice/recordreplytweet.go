package tweetservice

import (
	"cloud.google.com/go/firestore"
	"context"
	"errors"
	"fmt"
	"log"
	"mrkresnofatihdev/apps/gotwittercloneapp/application"
	"mrkresnofatihdev/apps/gotwittercloneapp/models"
	"strings"
)

const tweetReplyRecordCollectionName = "replyRecords"
const tweetReplyRecordStartsWithKeyFormat = "replies-user-starts-with#%s"

func RecordReplyTweet(ctx context.Context, replyTweet models.Tweet) {
	replyRecordExists, err := getTweetReplyRecordExists(ctx, replyTweet)
	if err != nil {
		log.Println("get-tweet-record-exists failed")
		return
	}

	if replyRecordExists {
		log.Println("reply record exists, will try to populate the found reply list")
		err = populateReplyList(ctx, replyTweet)
		if err != nil {
			log.Println("failed to populate reply list")
			return
		}
		log.Println("populated reply list")
	} else {
		log.Println("reply record does not exist, will try to init")
		err = initTweetReplyRecord(ctx, replyTweet)
		if err != nil {
			log.Println("failed to init reply record")
			return
		}
		log.Println("initialized new reply record")
	}
}

func initTweetReplyRecord(ctx context.Context, replyTweet models.Tweet) error {
	fireStr := application.GetFirestoreInstance()

	userTweetsCollectionKey := fmt.Sprintf(tweetsCollectionKeyFormat, replyTweet.ParentTweet.Username)
	tweetKey := replyTweet.ParentTweet.TweetId

	tweetReplyRecordKey := fmt.Sprintf(tweetReplyRecordStartsWithKeyFormat, replyTweet.Username[:1])

	newReplyRecord := models.TweetReplyRecordModel{
		UsernameStartsWith: replyTweet.Username[:1],
		Replies: map[string]string{
			replyTweet.TweetId: replyTweet.Username,
		},
	}

	_, err := fireStr.
		Collection(userTweetsCollectionKey).
		Doc(tweetKey).
		Collection(tweetReplyRecordCollectionName).
		Doc(tweetReplyRecordKey).
		Set(ctx, newReplyRecord)
	if err != nil {
		log.Println("error saving reply record to db")
		return errors.New("error saving reply record to db")
	}
	return nil
}

func getTweetReplyRecordExists(ctx context.Context, replyTweet models.Tweet) (bool, error) {
	fireStr := application.GetFirestoreInstance()

	userTweetsCollectionKey := fmt.Sprintf(tweetsCollectionKeyFormat, replyTweet.ParentTweet.Username)
	tweetKey := replyTweet.ParentTweet.TweetId

	tweetReplyRecordKey := fmt.Sprintf(tweetReplyRecordStartsWithKeyFormat, replyTweet.Username[:1])

	record, err := fireStr.
		Collection(userTweetsCollectionKey).
		Doc(tweetKey).
		Collection(tweetReplyRecordCollectionName).
		Doc(tweetReplyRecordKey).
		Get(ctx)
	if err != nil {
		if strings.HasPrefix(err.Error(), "rpc error: code = NotFound") {
			log.Println("not found error occurred")
			return false, nil
		}
		return false, err
	}

	return record.Exists(), nil
}

func populateReplyList(ctx context.Context, replyTweet models.Tweet) error {
	fireStr := application.GetFirestoreInstance()

	userTweetsCollectionKey := fmt.Sprintf(tweetsCollectionKeyFormat, replyTweet.ParentTweet.Username)
	tweetKey := replyTweet.ParentTweet.TweetId

	tweetReplyRecordKey := fmt.Sprintf(tweetReplyRecordStartsWithKeyFormat, replyTweet.Username[:1])

	_, err := fireStr.
		Collection(userTweetsCollectionKey).
		Doc(tweetKey).
		Collection(tweetReplyRecordCollectionName).
		Doc(tweetReplyRecordKey).
		Update(ctx, []firestore.Update{
			{
				Path:  fmt.Sprintf("replies.%s", replyTweet.TweetId),
				Value: replyTweet.Username,
			},
		})
	if err != nil {
		log.Println("failed to add reply to reply record")
		log.Println(err)
		return err
	}
	return nil
}
