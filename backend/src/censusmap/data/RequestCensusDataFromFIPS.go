package data

import (
  "net/http"
  "log"
  "io/ioutil"
)

const CensusAPI string = "http://api.census.gov"
const ACS52011 string = "/data/2011/acs5"
const CensusAPIKey string = "ab995d23ecc36a2920db2262f9ea8a9003ab2098"

func RequestCensusDataFromCodes(locationCodes CensusLocationCodes) []byte {
  values := "get=B17001_004E,B01001_026E" +
    "&for=block+group:" + string(locationCodes.BlockGroupCode) +
    "&in=state:" + string(locationCodes.StateCode) +
      "+county:" + string(locationCodes.CountyCode) +
      "+tract:" + string(locationCodes.TractCode) +
    "&key=" + CensusAPIKey
  res, err := http.Get(CensusAPI + ACS52011 + "?" + values)
  if err != nil {
    log.Fatal(err)
  }
  content, err := ioutil.ReadAll(res.Body)
  defer res.Body.Close()
  if err != nil {
    log.Fatal(err)
  }
  return content
}