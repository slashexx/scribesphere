package handlers

import (
	"encoding/json"
	"log"
	"fmt"
	"net/http"
	"scribesphere/models"
	"google.golang.org/api/iterator"
	"firebase.google.com/go/firestore"
	"scribesphere/firebaseSetup"

)

func CreatePost(w http.ResponseWriter, r *http.Request) {

	var newPost models.BlogPost

	json.NewDecoder(r.Body).Decode(&newPost)

	ctx := r.Context()

	_, _, err := firebaseSetup.Client.Collection("posts").Add(ctx, newPost)

	if err != nil {
		log.Fatalf("Failed adding post: %v", err)
		http.Error(w, "Failed to create post", http.StatusInternalServerError)
		return
	}


	w.WriteHeader(http.StatusCreated)
	if err := json.NewEncoder(w).Encode(newPost); err != nil {
		log.Fatalf("Failed to send response: %v", err)
	}
}
