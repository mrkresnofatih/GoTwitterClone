package timelineservice

import (
	"cloud.google.com/go/firestore"
	"context"
	"errors"
	"fmt"
	"github.com/google/uuid"
	"google.golang.org/api/iterator"
	"log"
	"mrkresnofatihdev/apps/gotwittercloneapp/application"
	"mrkresnofatihdev/apps/gotwittercloneapp/application/followservice"
	"mrkresnofatihdev/apps/gotwittercloneapp/application/playerservice"
	"mrkresnofatihdev/apps/gotwittercloneapp/models"
	"strconv"
	"strings"
	"time"
)

const PlayerHomeFeedQueryKeyFormat = "players-home-query#%s"
const PlayerHomeFeedBatchKeyFormat = "home-feed-batch#%s"
const HomeFeedCollectionName = "home-feeds"

func RunHomeFeedQuery(ctx context.Context, username string) error {
	err := setHomeFeedQueryToRunning(ctx, username)
	if err != nil {
		log.Println("set-home-feed-query-to-running failed")
		return err
	}

	now := time.Now().UnixMilli()
	currentStartsWith := "0"
	queryLimit := 10
	stop := false
	for !stop {
		query := models.FollowListQueryModel{
			Limit:      queryLimit,
			StartsWith: currentStartsWith,
			Username:   username,
		}
		followingLists, err := followservice.GetFollowingLists(ctx, query)
		if err != nil {
			log.Println("error executing get-following-lists")
			break
		}

		homeFeedBatchTweets := map[string]models.Tweet{}
		log.Println("followingLists is " + strconv.Itoa(len(followingLists)))
		for _, followList := range followingLists {
			for userName := range followList.FollowingList {
				log.Println(userName)
				profileQuery := models.ProfileFeedQueryModel{
					Username: userName,
					Limit:    4,
					StartAt:  now - 120000,
				}
				tweets, err := GetProfileFeed(ctx, profileQuery)
				if err != nil {
					log.Println("error get-profile-feed")
					continue
				}

				log.Println("tweets is " + strconv.Itoa(len(tweets)))
				for _, tweet := range tweets {
					log.Println("add new tweet to home-feed-batch-doc")
					homeFeedBatchTweets[tweet.TweetId] = tweet
				}
			}
			currentStartsWith = followList.StartsWith
		}

		if len(homeFeedBatchTweets) > 0 {
			log.Println("home-feed-batch-tweets found is not zero! will continue to save to home-feeds")
			err = writeHomeFeed(ctx, homeFeedBatchTweets, username)
			if err != nil {
				log.Println("something went wrong while writing home-feed")
			}
		} else {
			log.Println("no tweets found for new home-feed-batch doc! will not save to db!")
		}

		if len(followingLists) < queryLimit {
			stop = true
		}
	}

	return nil
}

func writeHomeFeed(ctx context.Context, tweets map[string]models.Tweet, username string) error {
	fireStr := application.GetFirestoreInstance()

	homeFeedKey := fmt.Sprintf(PlayerHomeFeedBatchKeyFormat, uuid.New().String())

	newHomeFeedBatchDoc := models.HomeFeedBatchTweets{
		Username:    username,
		Tweets:      tweets,
		CollectedAt: fmt.Sprintf("%015d", time.Now().UnixMilli()),
	}

	_, err := fireStr.
		Collection(HomeFeedCollectionName).
		Doc(homeFeedKey).
		Set(ctx, newHomeFeedBatchDoc)
	if err != nil {
		log.Println("something went wrong while saving home-feed-batch-doc")
		return errors.New("something went wrong while saving home-feed-batch-doc")
	}
	return nil
}

func GetHomeFeed(ctx context.Context, query models.HomeFeedQueryModel, username string) ([]models.HomeFeedBatchTweets, error) {
	returnDefaultOnError := func(err error) ([]models.HomeFeedBatchTweets, error) {
		return *new([]models.HomeFeedBatchTweets), err
	}

	fireStr := application.GetFirestoreInstance()

	var homeBatchTweets []models.HomeFeedBatchTweets

	log.Println("get iter get-home-feed")

	startAt := fmt.Sprintf("%015d", query.CollectedAt)
	maxLimit := 10
	if query.Limit > maxLimit {
		log.Println("query-limit get-home-feed too high! will default to max-limit")
		query.Limit = maxLimit
	}
	iter := fireStr.
		Collection(HomeFeedCollectionName).
		Where("username", "==", username).
		Where("collectedAt", ">", startAt).
		OrderBy("collectedAt", firestore.Desc).
		Limit(query.Limit).
		Documents(ctx)
	for {
		doc, err := iter.Next()
		if err == iterator.Done {
			log.Println("iter get-home-feed done")
			break
		}
		if err != nil {
			log.Println("error at iter " + err.Error())
			return returnDefaultOnError(err)
		}
		var batchTweet models.HomeFeedBatchTweets
		err = doc.DataTo(&batchTweet)
		if err != nil {
			log.Println("error cannot parse doc to tweet")
			continue
		}

		homeBatchTweets = append(homeBatchTweets, batchTweet)
	}
	return homeBatchTweets, nil
}

func setHomeFeedQueryToRunning(ctx context.Context, username string) error {
	playerCollectionKey := fmt.Sprintf(playerservice.PlayerKeyFormat, username)
	playerHomeFeedQueryKey := fmt.Sprintf(PlayerHomeFeedQueryKeyFormat, username)

	fireStr := application.GetFirestoreInstance()

	_, err := fireStr.
		Collection(playerCollectionKey).
		Doc(playerHomeFeedQueryKey).
		Update(ctx, []firestore.Update{
			{
				Path:  "collectedAt",
				Value: time.Now().UnixMilli(),
			},
		})
	if err != nil {
		log.Println("set-home-feed-query-to-running failed!")
		return errors.New("set-home-feed-query-to-running failed")
	}
	return nil
}

func GetHomeFeedQueryRunning(ctx context.Context, username string) (bool, error) {
	homeFeedQueryExists, err := getHomeFeedQueryExists(ctx, username)
	if err != nil {
		log.Println("error in get-home-feed-query-exists!")
		return false, errors.New("error in get-home-feed-query-exists")
	}

	if !homeFeedQueryExists {
		_, err := initHomeFeedQuery(ctx, username)
		if err != nil {
			log.Println("init home-feed-query-record failed")
			return false, errors.New("init home-feed-query-record failed")
		}
		return false, nil
	}

	homeFeedQueryData, err := getHomeFeedQuery(ctx, username)
	if err != nil {
		log.Println("get-home-feed-query failed")
		return false, err
	}
	nextQueryTime := homeFeedQueryData.LastQueryTime + 60000
	return nextQueryTime > time.Now().UnixMilli(), nil
}

func getHomeFeedQuery(ctx context.Context, username string) (models.HomeFeedQueryRecord, error) {
	fireStr := application.GetFirestoreInstance()

	playerCollectionKey := fmt.Sprintf(playerservice.PlayerKeyFormat, username)
	playerHomeFeedQueryKey := fmt.Sprintf(PlayerHomeFeedQueryKeyFormat, username)

	homeFeedQuery, err := fireStr.
		Collection(playerCollectionKey).
		Doc(playerHomeFeedQueryKey).
		Get(ctx)
	if err != nil {
		return *new(models.HomeFeedQueryRecord), err
	}

	var homeFeedQueryData models.HomeFeedQueryRecord
	err = homeFeedQuery.DataTo(&homeFeedQueryData)
	if err != nil {
		log.Println("parse to home-feed-query failed")
		return *new(models.HomeFeedQueryRecord), err
	}

	return homeFeedQueryData, nil
}

func getHomeFeedQueryExists(ctx context.Context, username string) (bool, error) {
	fireStr := application.GetFirestoreInstance()

	playerCollectionKey := fmt.Sprintf(playerservice.PlayerKeyFormat, username)
	playerHomeFeedQueryKey := fmt.Sprintf(PlayerHomeFeedQueryKeyFormat, username)

	homeFeedQuery, err := fireStr.
		Collection(playerCollectionKey).
		Doc(playerHomeFeedQueryKey).
		Get(ctx)
	if err != nil {
		if strings.HasPrefix(err.Error(), "rpc error: code = NotFound") {
			log.Println("home-feed-query record not found")
			return false, nil
		}
		return false, err
	}

	return homeFeedQuery.Exists(), nil
}

func initHomeFeedQuery(ctx context.Context, username string) (models.HomeFeedQueryRecord, error) {
	returnDefaultsOnError := func(err error) (models.HomeFeedQueryRecord, error) {
		return *new(models.HomeFeedQueryRecord), err
	}

	fireStr := application.GetFirestoreInstance()

	playerCollectionKey := fmt.Sprintf(playerservice.PlayerKeyFormat, username)
	playerHomeFeedQueryKey := fmt.Sprintf(PlayerHomeFeedQueryKeyFormat, username)

	homeFeedQuery := models.HomeFeedQueryRecord{
		Username:      username,
		LastQueryTime: 0,
	}

	_, err := fireStr.
		Collection(playerCollectionKey).
		Doc(playerHomeFeedQueryKey).
		Set(ctx, homeFeedQuery)
	if err != nil {
		return returnDefaultsOnError(err)
	}

	return homeFeedQuery, nil
}
