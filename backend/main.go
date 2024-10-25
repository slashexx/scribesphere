package main

import (
    "log"
    "net/http"
    "github.com/graphql-go/graphql"
)

func main() {
    http.HandleFunc("/graphql", func(w http.ResponseWriter, r *http.Request) {
        // Handle GraphQL requests
    })

    log.Println("Server is running on :8080")
    if err := http.ListenAndServe(":8080", nil); err != nil {
        log.Fatal(err)
    }
}
