package data

import (
  "io"
  "encoding/json"
  "log"
)

type PlainValueReport struct {
  BaseReport
  parsedConfig PlainValueConfigFormat
}

type PlainValueConfigFormat struct {
  BaseConfigFormat
  Vars map[string]string
}

func (r *PlainValueReport) ParseConfig(config []byte) {
  err := json.Unmarshal(config, &r.parsedConfig)
  if err != nil {
    log.Fatal(err)
  }
  r.requiredVariables = make([]string, 0)
  for _, code := range r.parsedConfig.Vars {
    r.requiredVariables = append(r.requiredVariables, code)
  }
}

func (r *PlainValueReport) GetRequiredVariables() []string {
  return r.requiredVariables
}

func (r *PlainValueReport) WriteFormattedReport(w io.Writer) {
  response := new(PlainValueConfigFormat)
  response.Kind = "plain_value"
  response.Vars = map[string]string{}
  for name, code := range r.parsedConfig.Vars {
    response.Vars[name] = r.variableValues[code]
  }
  encoded, err := json.Marshal(response)
  if err != nil {
    log.Fatal(err)
  }
  w.Write(encoded)
}
