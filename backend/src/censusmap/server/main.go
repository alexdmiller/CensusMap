package main

import (
    "fmt"
    "net/http"
    "censusmap/data"
)

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Printf("handler\n")
  location, codes := data.RequestLocationFromCoords(47.598755, -122.332764)
  fmt.Fprintf(w, "%v", location)
  fmt.Fprintf(w, "%v", codes)
}

func main() {
    http.HandleFunc("/", handler)
    http.ListenAndServe(":8080", nil)
}