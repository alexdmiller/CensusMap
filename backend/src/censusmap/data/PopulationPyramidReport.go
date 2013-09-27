package data

//import ("log")

type PopulationPyramidReport struct {
  BaseReport
  name string
}

type PopulationPyramidConfigFormat struct {
  Kind string `json:"kind"`
  Name string `json:"name"`
  Ages [][]string `json:"ages"`
}

func (r *PopulationPyramidReport) ParseConfig(config map[string]interface{}) {
  r.parsedConfig = config
  r.name = r.parsedConfig["name"].(string)
  ages := r.parsedConfig["ages"].([]interface{})
  for _, maleFemale := range ages {
    a := maleFemale.([]interface{})
    r.requiredVariables = append(r.requiredVariables, a[1].(string))
    r.requiredVariables = append(r.requiredVariables, a[2].(string))
  }
}

func (r *PopulationPyramidReport) getRequiredVariables() []string {
  return r.requiredVariables
}

func (r *PopulationPyramidReport) RequestAndParseData(codes CensusLocationCodes, key string) interface{} {
  variableValues := r.requestData(codes, key)
  response := new(PopulationPyramidConfigFormat)
  response.Kind = "population_pyramid"
  response.Name = r.name
  response.Ages = make([][]string, len(r.parsedConfig["ages"].([]interface{})))
  for i, maleFemaleCodes := range r.parsedConfig["ages"].([]interface{}) {
    a := maleFemaleCodes.([]interface{})
    maleResult := variableValues[a[1].(string)]
    femaleResult := variableValues[a[2].(string)]
    response.Ages[i] = []string{a[0].(string), maleResult, femaleResult}
  }
  return response
}
