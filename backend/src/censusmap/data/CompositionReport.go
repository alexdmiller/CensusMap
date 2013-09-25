package data

type CompositionReport struct {
  BaseReport
  name string
  sorted bool
  dropZeros bool
}

type CompositionConfigFormat struct {
  Kind string `json:"kind"`
  Total string `json:"total"`
  Name string `json:"name"`
  Parts [][]string `json:"parts"`
  Sorted bool `json:"sorted"`
  DropZeros bool `json:"dropZeros"`
}

func (r *CompositionReport) ParseConfig(config map[string]interface{}) {
  r.parsedConfig = config
  r.name = r.parsedConfig["name"].(string)
  r.sorted = r.parsedConfig["sorted"].(bool)
  r.dropZeros = r.parsedConfig["dropZeros"].(bool)
  r.requiredVariables = []string{r.parsedConfig["total"].(string)}
  parts := r.parsedConfig["parts"].([]interface{})
  for _, part := range parts {
    a := part.([]interface{})
    r.requiredVariables = append(r.requiredVariables, a[1].(string))
  }
}

func (r *CompositionReport) getRequiredVariables() []string {
  return r.requiredVariables
}

func (r *CompositionReport) RequestAndParseData(codes CensusLocationCodes) interface{} {
  variableValues := r.requestData(codes)
  response := new(CompositionConfigFormat)
  response.Kind = "composition"
  response.Name = r.name
  response.Sorted = r.sorted
  response.DropZeros = r.dropZeros
  response.Total = variableValues[r.parsedConfig["total"].(string)]
  response.Parts = make([][]string, len(r.parsedConfig["parts"].([]interface{})))
  for i, partResult := range r.parsedConfig["parts"].([]interface{}) {
    a := partResult.([]interface{})
    code := variableValues[a[1].(string)]
    response.Parts[i] = []string{a[0].(string), code}
  }
  return response
}
