package data

import (
  "testing"
)

var testFile = []byte(`[
  {
    "kind": "plain_value",
    "vars": {
      "Total Population": "B01003_001"
    }
  },
  {
    "kind": "composition",
    "total": "B02001_001",
    "name": "Racial Distribution",
    "parts": {
      "White": "B02001_002",
      "Black": "B02001_003",
      "Native American": "B02001_004",
      "Asian": "B02001_005",
      "Other": "B02001_007"
    }
  }
]`)

func TestParseConfigFile(t *testing.T) {
  r := new(CensusReports)
  r.ParseConfig(testFile)
}