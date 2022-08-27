package application

import (
	"context"
	firebase "firebase.google.com/go"
	"google.golang.org/api/option"
	"log"
)

var firebaseInstance *firebase.App

func GetFirebaseInstance() *firebase.App {
	if firebaseInstance == nil {
		ctx := context.Background()
		sa := option.WithCredentialsFile("firebase_conf.json")
		app, err := firebase.NewApp(ctx, nil, sa)
		if err != nil {
			log.Fatalln(err)
		}
		firebaseInstance = app
	}
	return firebaseInstance
}
