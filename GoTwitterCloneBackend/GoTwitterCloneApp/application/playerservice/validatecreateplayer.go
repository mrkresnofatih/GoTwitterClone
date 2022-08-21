package playerservice

import (
	"context"
	validator2 "github.com/go-playground/validator/v10"
	"mrkresnofatihdev/apps/gotwittercloneapp/models"
)

func ValidateCreatePlayer(ctx context.Context, createRequest models.PlayerCreateRequestModel) []string {
	validator := validator2.New()
	err := validator.Struct(createRequest)
	var listOfErrorMsg []string
	if err != nil {
		listOfErrors := err.(validator2.ValidationErrors)
		for _, er := range listOfErrors {
			listOfErrorMsg = append(listOfErrorMsg, er.Error())
		}
	}

	isPlayerExists := GetPlayerExists(ctx, createRequest.Username)
	if isPlayerExists {
		listOfErrorMsg = append(listOfErrorMsg, "player_username_exists")
	}
	return listOfErrorMsg
}
