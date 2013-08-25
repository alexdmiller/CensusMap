package data

import (
  "io"
  "encoding/json"
  "log"
  "fmt"
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

func (r *BaseReport) setVariable(name string, value string) {
  if (r.variableValues == nil) {
    r.variableValues = map[string]string{}
  }
  r.variableValues[name] = value
}

func (r *BaseReport) RequestData(codes CensusLocationCodes) {
  result := RequestCensusDataFromCodes(codes, []string{"B01003_001E", "B02001_001E"})
  fmt.Printf("%v", result)
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

