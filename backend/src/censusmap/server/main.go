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
var reports *data.CensusReports

func handler(w http.ResponseWriter, r *http.Request) {
  _, codes := data.RequestLocationFromCoords(47.598755, -122.332764)
  result := reports.RequestAndParseData(codes)
  resultJSON, err := json.Marshal(result)
  if err != nil {
    log.Fatal(err)
  }
  w.Write(resultJSON)
}

func main() {
  flag.StringVar(&configFileName, "c", "config/variable_codes.json", "path to configuration file")
  flag.Parse()
  config, err := ioutil.ReadFile(configFileName)
  if err != nil {
    log.Fatal(err)
  }
  reports = new(data.CensusReports)
  reports.ParseConfig(config)

  http.HandleFunc("/", handler)
  http.ListenAndServe(":8080", nil)
}