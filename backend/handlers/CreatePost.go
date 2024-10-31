package handlers

import (
	"encoding/json"
	"log"
	"net/http"
	"scribesphere/models"
	"scribesphere/firebaseSetup"
)

func CreatePost(w http.ResponseWriter, r *http.Request) {
	var newPost models.BlogPost

	// Decode the JSON request body into the newPost struct
	if err := json.NewDecoder(r.Body).Decode(&newPost); err != nil {
		http.Error(w, "Failed to decode JSON", http.StatusBadRequest)
		return
	}

	// Add the new post to Firestore
	ctx := r.Context() // Use the request context
	_, _, err := firebaseSetup.Client.Collection("posts").Add(ctx, newPost)
	if err != nil {
		log.Printf("Failed adding post: %v", err) // Log the error without terminating the server
		http.Error(w, "Failed to create post", http.StatusInternalServerError)
		return
	}

	// Send the response back to the client
	w.WriteHeader(http.StatusCreated)
	if err := json.NewEncoder(w).Encode(newPost); err != nil {
		log.Printf("Failed to send response: %v", err) // Log the error without terminating the server
		http.Error(w, "Failed to send response", http.StatusInternalServerError)
	}
}
