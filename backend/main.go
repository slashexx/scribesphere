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

	http.Handle("/posts", enableCORS(http.HandlerFunc(handlers.GetAllPosts)))
	http.Handle("/posts", enableCORS(http.HandlerFunc(handlers.CreatePost)))

	http.Handle("/post/", enableCORS(http.HandlerFunc(handlers.GetPost)))
	http.Handle("/post/", enableCORS(http.HandlerFunc(handlers.UpdatePost)))
	http.Handle("/post/", enableCORS(http.HandlerFunc(handlers.DeletePost)))

	fmt.Print("The server is now listening at port : 8080 (http://localhost:8080)")
	http.ListenAndServe(":8080", nil)

}
