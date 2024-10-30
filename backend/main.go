package main

import (
	"fmt"
	"net/http"
	"scribesphere/handlers"
	"scribesphere/models"
)

var posts []models.BlogPost

func main() {

	http.HandleFunc("/", func(written http.ResponseWriter, req *http.Request) {
		fmt.Fprintf(written, "Hello my niggas, checkout urls posts and post/id")
	})

	http.HandleFunc("/posts", func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case "GET":
			handlers.GetAllPosts(w, r)
		case "POST":
			handlers.CreatePost(w, r)
		default:
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		}
	})

	http.HandleFunc("/post/", func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case "GET":
			handlers.GetPost(w, r)
			break
		case "PUT":
			handlers.UpdatePost(w, r)
			break
		case "DELETE":
			handlers.DeletePost(w, r)

		default:
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)

		}
	})

	fmt.Print("The server is now listening at port : 8080 (http://localhost:8080)")
	http.ListenAndServe(":8080", nil)

}



