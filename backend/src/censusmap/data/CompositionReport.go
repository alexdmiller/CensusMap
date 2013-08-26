package data

type CompositionReport struct {
  BaseReport
  name string
}

type CompositionConfigFormat struct {
  Kind string `json:"kind"`
  Total string `json:"total"`
  Name string `json:"name"`
  Parts map[string]string `json:"parts"`
}

func (r *CompositionReport) ParseConfig(config map[string]interface{}) {
  r.parsedConfig = config
  r.name = r.parsedConfig["name"].(string)
  r.requiredVariables = []string{r.parsedConfig["total"].(string)}
  for _, code := range r.parsedConfig["parts"].(map[string]interface{}) {
    r.requiredVariables = append(r.requiredVariables, code.(string))
  }
}

func (r *CompositionReport) getRequiredVariables() []string {
  return r.requiredVariables
}

func (r *CompositionReport) RequestAndParseData(codes CensusLocationCodes) interface{} {
  variableValues := r.requestData(codes)
  response := new(CompositionConfigFormat)
  response.Kind = "composition"
  response.Parts = map[string]string{}
  for name, code := range r.parsedConfig["parts"].(map[string]interface{}) {
    response.Parts[name] = variableValues[code.(string)]
  }
  return response
}
