package playerservice

import (
	"context"
	validator2 "github.com/go-playground/validator/v10"
	"mrkresnofatihdev/apps/gotwittercloneapp/models"
)

func ValidateCreatePlayer(ctx context.Context, createRequest models.PlayerCreateRequestModel) []string {
	validator := validator2.New()
	err := validator.Struct(createRequest)
	if err != nil {
		var listOfErrorMsg []string
		listOfErrors := err.(validator2.ValidationErrors)
		for _, er := range listOfErrors {
			listOfErrorMsg = append(listOfErrorMsg, er.Error())
		}

		return listOfErrorMsg
	}
	return []string{}
}
