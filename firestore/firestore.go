package firestore

import (
	"context"
	"log"

	firebase "firebase.google.com/go/v4"
	// "github.com/googleapis/enterprise-certificate-proxy/client"
	"google.golang.org/api/option"
)

type Quote struct {
    PostContent string `firestore:"content"`
    Author      string `firestore:"author"`
    Title       string `firestore:"title"`
}

func InitializeFirebase() *firebase.App {
    sa := option.WithCredentialsFile("./firebaseConfig.json")
    app, err := firebase.NewApp(context.Background(), nil, sa)
    if err != nil {
        log.Fatalf("error initializing firebase app: %v\n", err)
    }
    return app
}

func AddQuote(app *firebase.App, quote *Quote) error {
    ctx := context.Background()
    client, err := app.Firestore(ctx)
    if err != nil {
		return err
    }
    defer client.Close()

    _, err = client.Collection("posts").NewDoc().Create(ctx, quote)
    return err
}

func GetQuote() *Quote {
	return &Quote{
		PostContent: "Fuc u",
		Author:      "Dhruv Puri",
		Title:       "Love & Life",
	}
}

// func main() {
// 	sa := option.WithCredentialsFile("./firebaseConfig.json")
// 	app, err := firebase.NewApp(context.Background(), nil, sa)

// 	client, err := app.Firestore(context.Background())

// 	if err != nil{
// 		log.Fatalln(err)
// 	}
// 	post := getQuote()
// 	log.Println(post)
	
// 	result, err := client.Collection("posts").NewDoc().Create(context.Background(), post)
// 	if err != nil {
// 		log.Fatalln(err)
// 	}

// 	log.Println(result)
// 	defer client.Close()

// }


