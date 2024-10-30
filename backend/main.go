package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)


var posts []BlogPost

func main() {

	http.HandleFunc("/", func(written http.ResponseWriter, req *http.Request) {
		fmt.Fprintf(written, "Hello my niggas, checkout urls posts and post/id")
	})

	http.HandleFunc("/posts", func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case "GET":
			GetAllPosts(w, r)
		case "POST":
			createPost(w, r)
		default:
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		}
	})

	http.HandleFunc("/post/", func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case "GET":
			getPost(w, r)
			break
		case "PUT":
			updatePost(w, r)
			break
		case "DELETE":
			deletePost(w, r)

		default:
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)

		}
	})

	fmt.Print("The server is now listening at port : 8080 (http://localhost:8080)")
	http.ListenAndServe(":8080", nil)

}


func getPost(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Path[len("/post/"):]
	for i := 0; i < len(posts); i++ {
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
	newPost.ID = fmt.Sprintf("%d", len(posts)+1)
	posts = append(posts, newPost)
	w.WriteHeader(http.StatusCreated)

	json.NewEncoder(w).Encode(newPost)
}

func updatePost(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Path[len("/post/"):]

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
	id := r.URL.Path[len("/post/"):]

	for i, post := range posts {
		if post.ID == id {
			posts = append(posts[:i], posts[i+1:]...)
			w.WriteHeader(http.StatusNoContent)
		}
	}

	http.Error(w, "Post to be deleted not found", http.StatusNotFound)
}
