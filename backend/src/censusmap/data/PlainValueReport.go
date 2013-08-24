package data

import (
  "io"
  "encoding/json"
  "log"
)

type PlainValueReport struct {
  BaseReport
}

type PlainValueConfigFormat struct {
  Vars map[string]string
}

func (r PlainValueReport) ParseConfig(config []byte) {
  parsed := new(PlainValueConfigFormat)
  err := json.Unmarshal(config, &parsed)
  if err != nil {
    log.Fatal(err)
  }
  r.requiredVariables = make([]string, 0)
  for _, code := range parsed.Vars {
    r.requiredVariables = append(r.requiredVariables, code)
  }
  log.Printf("%v", r.requiredVariables)
}

func (r PlainValueReport) WriteFormattedReport(w io.Writer) {
  
}
