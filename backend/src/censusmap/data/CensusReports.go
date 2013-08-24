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

import ("io"
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

func (r *BaseReport) SetVariable(name string, value string) {
  if (r.variableValues == nil) {
    panic("Must call ParseConfig before SetVariable")
  }
  r.variableValues[name] = value
}


type CensusReports struct {
  reports []Report

  requiredVariables map[string]bool
}

func (r *CensusReports) ParseConfig(config string) {

}

func (r *CensusReports) GetRequiredVariables() []string {
  return nil
}

func (r *CensusReports) WriteFormattedReports(w io.Writer) {

}

