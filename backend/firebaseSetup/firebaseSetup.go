package firebaseSetup

import (
	"context"
	firebase "firebase.google.com/go"
  "cloud.google.com/go/firestore"

	"google.golang.org/api/option"
	"log"
)

// Global variable to hold the Firestore client
var Client *firestore.Client

func InitialiseFirebase() {
	// Use a service account
	ctx := context.Background()
	sa := option.WithCredentialsFile("firebaseSetup.json")
	App, err := firebase.NewApp(ctx, nil, sa)
	if err != nil {
		log.Fatalln(err)
	}

	Client, err = App.Firestore(ctx)
	if err != nil {
		log.Fatalln(err)
	}
}
