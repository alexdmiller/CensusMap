package data

import (
  "testing"
  "fmt"
)

var testFile = []byte(`[
  {
    "kind": "plain_value",
    "vars": {
      "Total Population": "B01003_001E",
      "Other": "B02001_007E"
    }
  },
  {
    "kind": "plain_value",
    "vars": {
      "Total Population": "B01003_001E"
    }
  },
  {
    "kind": "composition",
    "total": "B02001_001E",
    "name": "Racial Distribution",
    "parts": {
      "White": "B02001_002E",
      "Black": "B02001_003E",
      "Native American": "B02001_004E",
      "Asian": "B02001_005E",
      "Other": "B02001_007E"
    }
  }
]`)

func TestCensusReportsParseConfig(t *testing.T) {
  r := new(CensusReports)
  r.ParseConfig(testFile)
  fmt.Printf("%v\n", len(r.reports))
}

func TestParseAndRequestData(t *testing.T) {
  _, codes := RequestLocationFromCoords(47.598755, -122.332764)
  r := new(CensusReports)
  r.ParseConfig(testFile)
  r.MakeRequests(codes)
}