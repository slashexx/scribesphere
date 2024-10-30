package firebasesetup

import (
	"log"
	"context"
	firebase "firebase.google.com/go"
	"google.golang.org/api/option"
  )
  
  // Use a service account
  ctx := context.Background()
  sa := option.WithCredentialsFile("firebaseSetup.json")
  app, err := firebase.NewApp(ctx, nil, sa)
  if err != nil {
	log.Fatalln(err)
  } 
  
  client, err := app.Firestore(ctx)
  if err != nil {
	log.Fatalln(err)
  }
  defer client.Close()