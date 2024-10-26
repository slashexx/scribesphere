package main

import (
	"context" // Import the context package
	"log"
	"net/http"

	"github.com/graphql-go/handler"
	firestore "scribesphere-backend/firestore"
)

func main() {
	app := firestore.InitializeFirebase() // Make sure to handle the initialization error
	if app == nil {
		log.Fatalf("Failed to initialize Firebase: %v", app)
	}

	h := handler.New(&handler.Config{
		Schema:   &schema,
		Pretty:   true,
		GraphiQL: true,
	})

	http.HandleFunc("/graphql", func(w http.ResponseWriter, r *http.Request) {
		ctx := context.WithValue(r.Context(), "firebaseApp", app)
		r = r.WithContext(ctx)
		h.ServeHTTP(w, r)
	})

	log.Println("Server is running on http://localhost:8080/graphql")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
