package data

import (
  "net/http"
  "log"
  "io/ioutil"
  "strings"
)

const CensusAPI string = "http://api.census.gov"
const ACS52011 string = "/data/2011/acs5"

/**
 * Uses the census.gov API to look up US Census Data for the passed Census location
 * and variables. Census variables can be found at:
 *  http://www.census.gov/developers/data/acs_5yr_2011_var.xml
 *
 * Currently uses American Community Survey data for 2011.
 * TODO: Parameter to determine which dataset is used.
 */
func RequestCensusDataFromCodes(key string, locationCodes CensusLocationCodes,
    variables []string) []byte {
  values := "get=" + strings.Join(variables, ",") +
    "&for=tract:" + string(locationCodes.TractCode) +
    "&in=state:" + string(locationCodes.StateCode) +
      "+county:" + string(locationCodes.CountyCode) +
    "&key=" + key
  res, err := http.Get(CensusAPI + ACS52011 + "?" + values)
  if err != nil {
    log.Println(err)
    panic("Could not retrieve data from " + CensusAPI + ". ")
  }
  content, err := ioutil.ReadAll(res.Body)
  defer res.Body.Close()
  if err != nil {
    log.Println(err)
    panic("Could not retrieve data from " + CensusAPI + ". ")
  }
  return content
}