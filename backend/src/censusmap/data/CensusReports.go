package data

import (
  "encoding/json"
  "log"
  "fmt"
)

type Report interface {
  ParseConfig(config map[string]interface{})
  RequestAndParseData(code CensusLocationCodes) interface{}
  requestData(code CensusLocationCodes) map[string]string
}

type BaseReport struct {
  requiredVariables []string
}

type BaseConfigFormat struct {
  Kind string `json:"kind"`
}

func (r *BaseReport) requestData(codes CensusLocationCodes) map[string]string {
  variableValues := map[string]string{}
  result := RequestCensusDataFromCodes(codes, r.requiredVariables)
  resultJSON := [][]string{}
  err := json.Unmarshal(result, &resultJSON)
  if err != nil {
    log.Printf("%s", result)
    log.Fatal(err)
  }
  for i := range resultJSON[0] {
    variableValues[resultJSON[0][i]] = resultJSON[1][i]
  }
  return variableValues
}


type CensusReports struct {
  reports []Report
  requiredVariables map[string]bool
}

func (r *CensusReports) ParseConfig(config []byte) {
  r.requiredVariables = map[string]bool{}
  var parsed []interface{}
  err := json.Unmarshal(config, &parsed)
  if err != nil {
    log.Fatal(err)
  }
  var report Report
  for i := range parsed {
    reportConfig := parsed[i].(map[string]interface{})
    kind := reportConfig["kind"].(string)
    switch kind {
    case "plain_value":
      report = new(PlainValueReport)
      report.ParseConfig(reportConfig)
    default:
      log.Print("Report kind " + kind + " not supported.")
    }
  }
}

func (r *CensusReports) MakeRequests(codes CensusLocationCodes) {
  for i := range r.reports {
    result := r.reports[i].RequestAndParseData(codes)
    fmt.Printf("%s\n", result)
  }
}