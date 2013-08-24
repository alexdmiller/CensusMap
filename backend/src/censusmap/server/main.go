package main

import (
  "net/http"
  "censusmap/data"
  "fmt"
)

func handler(w http.ResponseWriter, r *http.Request) {
  _, codes := data.RequestLocationFromCoords(47.598755, -122.332764)
  fmt.Fprintf(w, "%s", data.RequestCensusDataFromCodes(codes, []string{"B01003_001E", "B02001_001E"}))
}

func main() {
  http.HandleFunc("/", handler)
  http.ListenAndServe(":8080", nil)
}