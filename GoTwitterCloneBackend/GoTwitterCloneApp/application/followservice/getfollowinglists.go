package followservice

import (
	"cloud.google.com/go/firestore"
	"context"
	"google.golang.org/api/iterator"
	"log"
	"mrkresnofatihdev/apps/gotwittercloneapp/application"
	"mrkresnofatihdev/apps/gotwittercloneapp/models"
)

func GetFollowingLists(ctx context.Context, query models.FollowListQueryModel) ([]models.FollowingList, error) {
	returnDefaultOnError := func(err error) ([]models.FollowingList, error) {
		return *new([]models.FollowingList), err
	}

	fireStr := application.GetFirestoreInstance()

	var followingLists []models.FollowingList

	log.Println("get iter get-following-lists")
	maxLimit := 10
	if query.Limit > maxLimit {
		log.Println("query-limit get-following-list too high! will default to max-limit")
		query.Limit = maxLimit
	}
	iter := fireStr.
		Collection(followingListCollectionName).
		Where("username", "==", query.Username).
		Where("startsWith", ">", query.StartsWith).
		OrderBy("startsWith", firestore.Desc).
		Limit(query.Limit).
		Documents(ctx)
	for {
		doc, err := iter.Next()
		if err == iterator.Done {
			log.Println("iter get-following-lists done")
			break
		}
		if err != nil {
			log.Println("error at iter: " + err.Error())
			return returnDefaultOnError(err)
		}
		var followList models.FollowingList
		err = doc.DataTo(&followList)
		if err != nil {
			log.Println("error cannot parse doc to follow-list")
			continue
		}

		followingLists = append(followingLists, followList)
	}

	return followingLists, nil
}
