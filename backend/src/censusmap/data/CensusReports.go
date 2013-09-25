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

func (r *BaseReport) requestData(codes CensusLocationCodes) map[string]string {
  variableValues := map[string]string{}
  result := RequestCensusDataFromCodes(codes, r.requiredVariables)
  resultJSON := [][]string{}
  err := json.Unmarshal(result, &resultJSON)
  if err != nil {
    log.Printf("%s", result)
    panic("JSON parsing error: " + string(result))
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
    log.Printf("Error when parsing ")
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
    case "population_pyramid":
      report = new(PopulationPyramidReport)
    default:
      log.Print("Report kind '" + kind + "' not supported.")
    }
    report.ParseConfig(reportConfig)
    r.reports = append(r.reports, report)
  }
}

type ReportAndPosition struct {
  Position int
  Report interface{}
}

func (r *CensusReports) RequestAndParseData(codes CensusLocationCodes) []interface{} {
  var wg sync.WaitGroup
  wg.Add(len(r.reports))
  ch := make(chan ReportAndPosition)
  for i := range r.reports {
    go func(report Report, ch chan ReportAndPosition, wg *sync.WaitGroup, position int) {
      defer wg.Done()
      reportResult := report.RequestAndParseData(codes).(interface{})
      ch <- ReportAndPosition{position, reportResult}
    }(r.reports[i], ch, &wg, i)
  }

  go func() {
    wg.Wait()
    close(ch)
  }()

  results := make([]interface{}, len(r.reports), len(r.reports))
  for result := range ch {
    results[result.Position] = result.Report
  }

  return results
}