package data

import (
	"net/http"
	"net/url"
	"io/ioutil"
	"strconv"
	"log"
)

func RequestLocationFromCoords(lat float64, lon float64) []byte {
	stringParams := []string{
		strconv.FormatFloat(lat, 'b', -1, 64),
		strconv.FormatFloat(lon, 'b', -1, 64),
	}
	values := url.Values{
		"format": {"json"},
		"latitude": stringParams[0:1],
		"longitude": stringParams[1:2],
		"showall": {"true"},
	}
	res, err := http.Get("http://data.fcc.gov/api/block/find" + values.Encode())
	if err != nil {
		log.Fatal(err)
	}
	content, err := ioutil.ReadAll(res.Body)
	res.Body.Close()
	if err != nil {
		log.Fatal(err)
	}
	return content
}