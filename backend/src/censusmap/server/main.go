package main

import (
    "fmt"
    "net/http"
    "censusmap/data"
)

func handler(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintf(w, "%v", data.RequestLocationFromCoords(47.598755, -122.332764))
}

func main() {
    http.HandleFunc("/", handler)
    http.ListenAndServe(":8080", nil)
}

