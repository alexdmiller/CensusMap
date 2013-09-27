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
var key string
var reports *data.CensusReports

// TODO: return error if request malformatted
func handler(w http.ResponseWriter, r *http.Request) {
  r.ParseForm()
  log.Printf("Request: %s, %s", r.Form["lat"][0], r.Form["long"][0])
  defer func(w http.ResponseWriter) {
    if r := recover(); r != nil {
      log.Println("Error: ", r)
      w.Write([]byte(r.(string)))
    }
  }(w)
  log.Printf("Requesting census tract code from FCC")
  location, codes := data.RequestLocationFromCoords(r.Form["lat"][0], r.Form["long"][0])
  log.Printf("Requesting census data from census.gov")
  reportResults := reports.RequestAndParseData(codes, key)
  result := map[string]interface{}{}
  result["reports"] = reportResults
  result["tract"] = string(codes.TractCode)
  result["county"] = location.County
  result["state"] = location.State
  resultJSON, err := json.Marshal(result)
  if err != nil {
    log.Printf("%s", err)
    w.Write(resultJSON)
  } else {
    w.Write(resultJSON)
    log.Printf("Response sent for %s, %s", r.Form["lat"][0], r.Form["long"][0])  
  }
}

func main() {
  flag.StringVar(&configFileName, "c", "config/variable_codes.json", "path to configuration file")
  flag.StringVar(&wwwDirectory, "w", "/tmp", "path to www directory")
  flag.StringVar(&key, "k", "", "census.gov API key")
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