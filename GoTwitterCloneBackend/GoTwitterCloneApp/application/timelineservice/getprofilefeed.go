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

func GetProfileFeed(ctx context.Context, query models.ProfileFeedQueryModel) ([]models.Tweet, error) {
	returnDefaultOnError := func(err error) ([]models.Tweet, error) {
		return *new([]models.Tweet), err
	}

	fireStr := application.GetFirestoreInstance()

	var tweets []models.Tweet

	log.Println("get iter")

	startAt := fmt.Sprintf("%015d", query.StartAt)
	iter := fireStr.
		Collection(tweetservice.TweetsCollectionName).
		Where("username", "==", query.Username).
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
		var tweet models.Tweet
		err = doc.DataTo(&tweet)
		if err != nil {
			log.Println("error cannot parse doc to tweet")
			continue
		}

		tweets = append(tweets, tweet)
	}

	return tweets, nil
}
