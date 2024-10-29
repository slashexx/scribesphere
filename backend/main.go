package main

import (
	"context"
	"encoding/json"
	"log"
	"net/http"
	firestore "scribesphere-backend/firestore"
)

func main() {
	// Initialize Firebase
	app := firestore.InitializeFirebase() // Ensure you handle initialization error
	if app == nil {
		log.Fatalf("Failed to initialize Firebase: %v", app)
	}

	// Define REST API routes
	http.HandleFunc("/articles", func(w http.ResponseWriter, r *http.Request) {
		ctx := context.WithValue(r.Context(), "firebaseApp", app)
		switch r.Method {
		case http.MethodGet:
			getArticles(w, r.WithContext(ctx))
		case http.MethodPost:
			createArticle(w, r.WithContext(ctx))
		default:
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		}
	})

	log.Println("Server is running on http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}

// Example handler for GET request (retrieving articles)
func getArticles(w http.ResponseWriter, r *http.Request) {
	// Replace with actual Firestore query to retrieve articles
	articles := []string{"Article 1", "Article 2"} // Sample data
	json.NewEncoder(w).Encode(articles)
}

// Example handler for POST request (creating an article)
func createArticle(w http.ResponseWriter, r *http.Request) {
	// Example to parse JSON body for creating a new article
	var article struct {
		Title string `json:"title"`
		Body  string `json:"body"`
	}
	if err := json.NewDecoder(r.Body).Decode(&article); err != nil {
		http.Error(w, "Bad request", http.StatusBadRequest)
		return
	}

	// Replace with actual Firestore logic to save article
	log.Printf("Article created: %v", article)
	w.WriteHeader(http.StatusCreated)
}
