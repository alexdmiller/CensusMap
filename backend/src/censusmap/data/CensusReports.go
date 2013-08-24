package data
/*
parseConfig(configFile / stream)
getCensusCodes() string
  implement with set
formatData(data) stream?
  create empty 'data packs' from config
  loop over every entry in data
    find all 'data packs' interested in census code, give value to them
  loop over data packs
    ask pack to write itself to stream?0
*/

import (
  "io"
  "encoding/json"
  "log"
)

type Report interface {
  ParseConfig(config []byte)
  GetRequiredVariables() []string
  SetVariable(name string, value string)
  WriteFormattedReport(w io.Writer)
}

type BaseReport struct {
  requiredVariables []string
  variableValues map[string]string
}

type BaseConfigFormat struct {
  Kind string `json:"kind"`
}

func (r *BaseReport) SetVariable(name string, value string) {
  if (r.variableValues == nil) {
    r.variableValues = map[string]string{}
  }
  r.variableValues[name] = value
}


type CensusReports struct {
  reports []Report
  requiredVariables map[string]bool
}

func (r *CensusReports) ParseConfig(config []byte) {
  var parsed []interface{}
  err := json.Unmarshal(config, &parsed)
  if err != nil {
    log.Fatal(err)
  }
  for i := range parsed {
    log.Printf("%v", parsed[i])
  }
}

func (r *CensusReports) GetRequiredVariables() []string {
  return nil
}

func (r *CensusReports) WriteFormattedReports(w io.Writer) {

}

