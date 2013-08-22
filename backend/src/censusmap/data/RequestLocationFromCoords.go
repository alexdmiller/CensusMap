package data

import (
	"fmt"
	"net/http"
	"net/url"
	"io/ioutil"
	"strconv"
	"log"
	"encoding/json"
)

type CensusBlock struct {
	FIPS string
}

type LocationResponse struct {
	Block CensusBlock
}

func RequestLocationFromCoords(lat float64, lon float64) CensusBlock {
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
	reqURL := "http://data.fcc.gov/api/block/find?" + values.Encode()
	fmt.Printf(reqURL + "\n")
	res, err := http.Get("http://data.fcc.gov/api/block/find?" + values.Encode())
	if err != nil {
		log.Fatal(err)
	}
	content, err := ioutil.ReadAll(res.Body)
	fmt.Printf("%s\n", content)
	defer res.Body.Close()
	if err != nil {
		log.Fatal(err)
	}
	var location LocationResponse
	err = json.Unmarshal(content, &location)
	if err != nil {
		log.Fatal(err)
	}
	return location.Block
}