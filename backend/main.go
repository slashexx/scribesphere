package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type BlogPost struct {
	ID      string "json:id"
	Title   string "json:title"
	Content string "json:content"
	Author  string "json:author"
}

var posts []BlogPost

func main() {

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello my niggas")
	})

	fmt.Print("The server is now listening at port : 8080 (http://localhost:8080)")
	http.ListenAndServe(":8080", nil)

}

func getAllPosts(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(posts)
}

func getPost(w http.ResponseWriter, r *http.Request){
	id := r.URL.Path[len("/post/"):]
	for i := range len(posts){
		if posts[i].ID == id {
			json.NewEncoder(w).Encode(posts[i])
			return
		}
	}

	http.Error(w, "Specified post does not exist", http.StatusNotFound)
}

func createPost(w http.ResponseWriter, r *http.Request) {
	var newPost BlogPost
	json.NewDecoder(r.Body).Decode(&newPost)
	newPost.ID = fmt.Sprint("%d", len(posts)+1)
	posts = append(posts, newPost)
	w.WriteHeader(http.StatusCreated)

	json.NewEncoder(w).Encode(newPost)
}

func updatePost(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Path[len("/posts/"):]

	var updatedPost BlogPost
	json.NewDecoder(r.Body).Decode(&updatedPost)

	for i, post := range posts {
		if post.ID == id {
			posts[i] = updatedPost
			json.NewEncoder(w).Encode(posts[i])
			return
		}
	}

	http.Error(w, "Post not found", http.StatusNotFound)
}

func deletePost(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Path[len("/posts/"):]

	for i, post := range posts {
		if post.ID == id {
			posts = append(posts[:i], posts[i+1:]...)
			w.WriteHeader(http.StatusNoContent)
		}
	}

	http.Error(w, "Post to be deleted not found", http.StatusNotFound)
}
