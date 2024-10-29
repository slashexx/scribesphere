package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	firestore "scribesphere-backend/firestore" // Assuming Firestore package setup
)

// Article represents the structure of an article in the application
type Article struct {
	ID    string `json:"id,omitempty"`
	Title string `json:"title"`
	Body  string `json:"body"`
}

func main() {
	// Initialize Firebase app
	app := firestore.InitializeFirebase()
	if app == nil {
		log.Fatalf("Failed to initialize Firebase: %v", app)
	}

	// Define routes for the REST API
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

// getArticles handles GET requests to retrieve articles
func getArticles(w http.ResponseWriter, r *http.Request) {
	// Sample data for demonstration; replace with Firestore query logic
	articles := []Article{
		{ID: "1", Title: "First Article", Body: "This is the first article."},
		{ID: "2", Title: "Second Article", Body: "This is the second article."},
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(articles)
}

// createArticle handles POST requests to add a new article
func createArticle(w http.ResponseWriter, r *http.Request) {
	var article Article
	if err := json.NewDecoder(r.Body).Decode(&article); err != nil {
		http.Error(w, "Bad request", http.StatusBadRequest)
		return
	}

	// Placeholder for saving to Firestore
	article.ID = "generated_id" // Replace with actual Firestore ID generation
	fmt.Printf("Article created: %+v\n", article) // Log to console for verification

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(article) // Respond with the created article
}
