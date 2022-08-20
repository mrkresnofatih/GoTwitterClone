package application

import (
	"context"
	firebase "firebase.google.com/go"
	"google.golang.org/api/option"
	"log"
)

var firebaseInstance *firebase.App

var firebaseCtxInstance *context.Context

func getFirebaseContextInstance() *context.Context {
	if firebaseCtxInstance == nil {
		ctx := context.Background()
		firebaseCtxInstance = &ctx
	}
	return firebaseCtxInstance
}

func GetFirebaseInstance() *firebase.App {
	if firebaseInstance == nil {
		ctx := getFirebaseContextInstance()
		sa := option.WithCredentialsFile("firebase_conf.json")
		app, err := firebase.NewApp(*ctx, nil, sa)
		if err != nil {
			log.Fatalln(err)
		}
		firebaseInstance = app
	}
	return firebaseInstance
}
