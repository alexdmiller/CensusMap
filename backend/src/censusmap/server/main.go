package main

import (
    "fmt"
    "net/http"
)

func handler(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintf(w, "%s", r.URL.Path[1:])
}

func main() {
    http.HandleFunc("/", handler)
    fmt.Printf("asdfasdfasdf running the server...\n")
    http.ListenAndServe(":8080", nil)
}

