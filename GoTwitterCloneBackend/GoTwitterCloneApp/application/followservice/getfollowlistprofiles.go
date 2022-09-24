package followservice

import (
	"context"
	"log"
	"mrkresnofatihdev/apps/gotwittercloneapp/application/playerservice"
	"mrkresnofatihdev/apps/gotwittercloneapp/models"
)

func GetFollowerListProfiles(ctx context.Context, query models.FollowListQueryModel) ([]models.PlayerGetMinimumProfileResponseModel, error) {
	returnDefaultOnError := func(err error) ([]models.PlayerGetMinimumProfileResponseModel, error) {
		return *new([]models.PlayerGetMinimumProfileResponseModel), err
	}

	followerLists, err := GetFollowerLists(ctx, query)
	if err != nil {
		log.Println("get-follower-lists failed")
		return returnDefaultOnError(err)
	}

	var players []models.PlayerGetMinimumProfileResponseModel
	for _, followerList := range followerLists {
		for followerUsername := range followerList.FollowerList {
			player, err := playerservice.GetPlayerMinimumProfile(ctx, followerUsername)
			if err != nil {
				log.Println("get-player-minimum-profile not found")
				continue
			}
			if players == nil {
				players = []models.PlayerGetMinimumProfileResponseModel{
					player,
				}
			} else {
				players = append(players, player)
			}
		}
	}

	return players, nil
}

func GetFollowingListProfiles(ctx context.Context, query models.FollowListQueryModel) ([]models.PlayerGetMinimumProfileResponseModel, error) {
	returnDefaultOnError := func(err error) ([]models.PlayerGetMinimumProfileResponseModel, error) {
		return *new([]models.PlayerGetMinimumProfileResponseModel), err
	}

	followingLists, err := GetFollowingLists(ctx, query)
	if err != nil {
		log.Println("get-following-lists failed")
		return returnDefaultOnError(err)
	}

	var players []models.PlayerGetMinimumProfileResponseModel
	for _, followingList := range followingLists {
		for followingUsername := range followingList.FollowingList {
			player, err := playerservice.GetPlayerMinimumProfile(ctx, followingUsername)
			if err != nil {
				log.Println("get-player-minimum-profile not found")
				continue
			}
			if players == nil {
				players = []models.PlayerGetMinimumProfileResponseModel{
					player,
				}
			} else {
				players = append(players, player)
			}
		}
	}

	return players, nil
}
