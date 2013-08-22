package main

import (
    "fmt"
    "net/http"
    "censusmap/data"
)

func handler(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintf(w, "%v", data.RequestLocationFromCoords(50, 30.2))
}

func main() {
    http.HandleFunc("/", handler)
    http.ListenAndServe(":8080", nil)
}

