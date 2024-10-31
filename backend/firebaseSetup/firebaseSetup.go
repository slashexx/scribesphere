package firebaseSetup

import (
    "context"
    "log"
    firebase "firebase.google.com/go"
    "google.golang.org/api/option"
)

func InitializeFirebase() *firebase.App {
    
    opt := option.WithCredentialsFile("firebaseSetup/firebaseConfig.json")
    
    App, err := firebase.NewApp(context.Background(), nil, opt)
    if err != nil {
        log.Fatalf("Firebase initialization error: %v\n", err)
    }
    
    return App
}
