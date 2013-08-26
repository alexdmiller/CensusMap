package main

import (
  "net/http"
  "censusmap/data"
  "encoding/json"
  "log"
  "io/ioutil"
  "flag"
)

var configFileName string
var wwwDirectory string
var reports *data.CensusReports

func handler(w http.ResponseWriter, r *http.Request) {
  r.ParseForm()
  log.Printf("Request: %s, %s", r.Form["lat"][0], r.Form["long"][0])
  // 47.598755, -122.332764
  _, codes := data.RequestLocationFromCoords(r.Form["lat"][0], r.Form["long"][0])
  result := reports.RequestAndParseData(codes)
  resultJSON, err := json.Marshal(result)
  if err != nil {
    log.Fatal(err)
  }
  w.Write(resultJSON)
}

func main() {
  flag.StringVar(&configFileName, "c", "config/variable_codes.json", "path to configuration file")
  flag.StringVar(&wwwDirectory, "w", "/tmp", "path to www directory")
  flag.Parse()
  config, err := ioutil.ReadFile(configFileName)
  if err != nil {
    log.Fatal(err)
  }
  reports = new(data.CensusReports)
  reports.ParseConfig(config)
  
  http.Handle("/", http.FileServer(http.Dir(wwwDirectory)))
  http.HandleFunc("/api/census", handler)
  http.ListenAndServe(":8080", nil)
}