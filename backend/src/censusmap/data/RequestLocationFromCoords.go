package data

import (
	"net/http"
	"net/url"
	"io/ioutil"
	"strconv"
	"log"
	"encoding/json"
)

const FCCBlockAPI string = "http://data.fcc.gov/api/block/find"

// For marshalling the response from fcc.gov
type FCCLocationResponse struct {
  County struct {
    FIPS, Name string
  }
  State struct {
    Code, FIPS, Name string
  }
	Block struct {
		FIPS string
	}
}

type CensusLocation struct {
	State, County string
}

// Stores FIPS identification codes for geographic levels
type CensusLocationCodes struct {
	StateCode, CountyCode, TractCode, BlockGroupCode []byte
}

/**
 * Uses an FCC API to find the block group FIPS code for the passed latitude and
 * longitude. A block group is the finest granularity of geographic data released
 * by the US Census.  
 */
func RequestLocationFromCoords(lat float64, lon float64) (CensusLocation, CensusLocationCodes) {
	stringParams := []string{
		strconv.FormatFloat(lat, 'f', -1, 64),
		strconv.FormatFloat(lon, 'f', -1, 64),
	}
	values := url.Values{
		"format": {"json"},
		"latitude": stringParams[0:1],
		"longitude": stringParams[1:2],
		"showall": {"true"},
	}
	res, err := http.Get(FCCBlockAPI + "?" + values.Encode())
	if err != nil {
		log.Fatal(err)
	}
	content, err := ioutil.ReadAll(res.Body)
	defer res.Body.Close()
	if err != nil {
		log.Fatal(err)
	}
	var locationResponse FCCLocationResponse
	err = json.Unmarshal(content, &locationResponse)
	if err != nil {
		log.Fatal(err)
	}
  location := CensusLocation{
    locationResponse.State.Name,
    locationResponse.County.Name,
  }
  bytes := []byte(locationResponse.Block.FIPS)
  locationCodes := CensusLocationCodes{
    bytes[0:2],
    bytes[2:5],
    bytes[5:11],
    bytes[11:12],
  }
	return location, locationCodes
}