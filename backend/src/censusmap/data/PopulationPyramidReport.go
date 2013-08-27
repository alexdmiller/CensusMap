package data

//import ("log")

type PopulationPyramidReport struct {
  BaseReport
  name string
}

type PopulationPyramidConfigFormat struct {
  Kind string `json:"kind"`
  Name string `json:"name"`
  Ages map[string][]string `json:"ages"`
}

func (r *PopulationPyramidReport) ParseConfig(config map[string]interface{}) {
  r.parsedConfig = config
  r.name = r.parsedConfig["name"].(string)
  ages := r.parsedConfig["ages"].(map[string]interface{})
  for _, maleFemale := range ages {
    a := maleFemale.([]interface{})
    r.requiredVariables = append(r.requiredVariables, a[0].(string))
    r.requiredVariables = append(r.requiredVariables, a[1].(string))
  }
}

func (r *PopulationPyramidReport) getRequiredVariables() []string {
  return r.requiredVariables
}

func (r *PopulationPyramidReport) RequestAndParseData(codes CensusLocationCodes) interface{} {
  variableValues := r.requestData(codes)
  response := new(PopulationPyramidConfigFormat)
  response.Kind = "population_pyramid"
  response.Name = r.name
  response.Ages = map[string][]string{}
  for name, maleFemaleCodes := range r.parsedConfig["ages"].(map[string][]string) {
    maleResult := variableValues[maleFemaleCodes[0]]
    femaleResult := variableValues[maleFemaleCodes[1]]
    response.Ages[name] = []string{maleResult, femaleResult}
  }
  return response
}
