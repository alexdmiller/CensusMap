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
  location, codes := data.RequestLocationFromCoords(r.Form["lat"][0], r.Form["long"][0])
  reportResults := reports.RequestAndParseData(codes)
  result := map[string]interface{}{}
  result["reports"] = reportResults
  result["tract"] = string(codes.TractCode)
  result["county"] = location.County
  result["state"] = location.State
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