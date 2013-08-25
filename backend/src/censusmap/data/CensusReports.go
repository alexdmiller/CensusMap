package data

import (
  "io"
  "encoding/json"
  "log"
)

type Report interface {
  ParseConfig(config map[string]interface{})
  RequestData(code CensusLocationCodes)
  WriteFormattedReport(w io.Writer)
}

type BaseReport struct {
  requiredVariables []string
  variableValues map[string]string
}

type BaseConfigFormat struct {
  Kind string `json:"kind"`
}

// TODO: Add constructor rather than check nil every call
func (r *BaseReport) setVariable(name string, value string) {
  if (r.variableValues == nil) {
    r.variableValues = map[string]string{}
  }
  r.variableValues[name] = value
}

func (r *BaseReport) RequestData(codes CensusLocationCodes) {
  result := RequestCensusDataFromCodes(codes, r.requiredVariables)
  resultJSON := [][]string{}
  err := json.Unmarshal(result, &resultJSON)
  if err != nil {
    log.Printf("%s", result)
    log.Fatal(err)
  }
  for i := range resultJSON[0] {
    r.setVariable(resultJSON[0][i], resultJSON[1][i])
  }
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

func (r *CensusReports) MakeRequests() {

}

func keys(m map[string]bool) []string {
  keys := []string{}
  for k := range m {
    keys = append(keys, k)
  }
  return keys
}

func (r *CensusReports) WriteFormattedReports(w io.Writer) {

}

