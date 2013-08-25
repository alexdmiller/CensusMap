package data

import (
  "testing"
  "fmt"
  "sort"
  "encoding/json"
)

const dummyConfig string = `{"kind": "plain_value", "vars": {"Total Population": "B01003_001E", "Other": "B02001_007E"}}`

func newPlainValueReport(t *testing.T) *PlainValueReport {
  r := new(PlainValueReport)
  bytes := []byte(dummyConfig)
  var parsed map[string]interface{}
  err := json.Unmarshal(bytes, &parsed)
  if err != nil {
    t.Error(err)
  }
  r.ParseConfig(parsed)
  return r
}

func TestParseConfig(t *testing.T) {
  r := newPlainValueReport(t)
  required := r.requiredVariables
  expected := []string{"B01003_001E", "B02001_007E"}
  sort.Strings(required)
  sort.Strings(expected)
  requiredString := fmt.Sprintf("%v", required)
  expectedString := fmt.Sprintf("%v", expected)
  if requiredString !=  expectedString {
    t.Errorf("Expected %v but got %v", expectedString, requiredString)
  }
}

func TestRequestAndParseData(t *testing.T) {
  _, codes := RequestLocationFromCoords(47.598755, -122.332764)
  r := newPlainValueReport(t)
  result := r.RequestAndParseData(codes).(*PlainValueConfigFormat)
  if result.Kind != "plain_value" {
    t.Errorf(`Kind should be "plain_value", but was %s.`, result.Kind)
  }
  if len(result.Vars) != 2 {
    t.Errorf("Vars should have 2 results, but instead has %d.", len(result.Vars))
  }
}