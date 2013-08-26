package data

import (
  "encoding/json"
  "log"
  "sync"
)

type Report interface {
  ParseConfig(config map[string]interface{})
  RequestAndParseData(code CensusLocationCodes) interface{}
  requestData(code CensusLocationCodes) map[string]string
}

type BaseReport struct {
  requiredVariables []string
  parsedConfig map[string]interface{}
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
  r.reports = []Report{}
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
    case "composition":
      report = new(CompositionReport)
    default:
      log.Print("Report kind '" + kind + "' not supported.")
    }
    report.ParseConfig(reportConfig)
      r.reports = append(r.reports, report)
  }
}

func (r *CensusReports) RequestAndParseData(codes CensusLocationCodes) []interface{} {
  var wg sync.WaitGroup
  wg.Add(len(r.reports))
  ch := make(chan interface{})
  for i := range r.reports {
    go func(report Report, ch chan interface{}, wg *sync.WaitGroup) {
      defer wg.Done()
      reportResult := report.RequestAndParseData(codes)
      ch <- reportResult.(interface{})
    }(r.reports[i], ch, &wg)
  }

  go func() {
    wg.Wait()
    close(ch)
  }()

  results := []interface{}{}
  for result := range ch {
    results = append(results, result)
  }

  return results
}