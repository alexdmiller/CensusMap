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
  Kind string
  Vars map[string]string
}

func (r PlainValueReport) ParseConfig(config []byte) {
  parsed := new(PlainValueConfigFormat)
  log.Printf("%s", config)
  err := json.Unmarshal(config, &parsed)
  if err != nil {
    log.Fatal(err)
  }
}

func (r PlainValueReport) WriteFormattedReport(w io.Writer) {
  
}
