package data

import (
  "io"
  "encoding/json"
  "log"
)

type PlainValueReport struct {
  BaseReport
  parsedConfig map[string]interface{}
}

type PlainValueConfigFormat struct {
  Kind string `json:"kind"`
  Vars map[string]string `json:"vars"`
}

func (r *PlainValueReport) ParseConfig(config map[string]interface{}) {
  r.parsedConfig = config
  r.requiredVariables = make([]string, 0)
  for _, code := range r.parsedConfig["vars"].(map[string]interface{}) {
    r.requiredVariables = append(r.requiredVariables, code.(string))
  }
}

func (r *PlainValueReport) GetRequiredVariables() []string {
  return r.requiredVariables
}

func (r *PlainValueReport) WriteFormattedReport(w io.Writer) {
  response := new(PlainValueConfigFormat)
  response.Kind = "plain_value"
  response.Vars = map[string]string{}
  for name, code := range r.parsedConfig["vars"].(map[string]interface{}) {
    response.Vars[name] = r.variableValues[code.(string)]
  }
  encoded, err := json.Marshal(response)
  if err != nil {
    log.Fatal(err)
  }
  w.Write(encoded)
}
