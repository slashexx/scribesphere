package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"scribesphere/models"
)

func CreatePost(w http.ResponseWriter, r *http.Request) {
	var newPost models.BlogPost
	json.NewDecoder(r.Body).Decode(&newPost)
	newPost.ID = fmt.Sprintf("%d", len(posts)+1)
	posts = append(posts, newPost)
	w.WriteHeader(http.StatusCreated)

	json.NewEncoder(w).Encode(newPost)
}
