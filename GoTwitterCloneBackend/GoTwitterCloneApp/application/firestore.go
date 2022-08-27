package application

import (
	"cloud.google.com/go/firestore"
	"context"
	"log"
)

var firestoreInstance *firestore.Client

func GetFirestoreInstance() *firestore.Client {
	if firestoreInstance == nil {
		firebaseApp := GetFirebaseInstance()
		firebaseCtxInstance := context.Background()
		client, err := firebaseApp.Firestore(firebaseCtxInstance)
		if err != nil {
			log.Fatalln(err)
		}
		firestoreInstance = client
	}
	return firestoreInstance
}

func CloseFirestore() {
	if firestoreInstance != nil {
		err := firestoreInstance.Close()
		if err != nil {
			return
		}
	}
}
