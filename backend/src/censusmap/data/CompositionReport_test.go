package data

import (
  "testing"
  "fmt"
  "sort"
  "encoding/json"
)

const compositionDummyConfig string = `{
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
  }`

func newCompositionReport(t *testing.T) *CompositionReport {
  r := new(CompositionReport)
  bytes := []byte(compositionDummyConfig)
  var parsed map[string]interface{}
  err := json.Unmarshal(bytes, &parsed)
  if err != nil {
    t.Error(err)
  }
  r.ParseConfig(parsed)
  return r
}

func TestParseConfigComposition(t *testing.T) {
  r := newCompositionReport(t)
  required := r.requiredVariables
  expected := []string{"B02001_001E", "B02001_002E", "B02001_003E", "B02001_004E", "B02001_005E", "B02001_007E"}
  sort.Strings(required)
  sort.Strings(expected)
  requiredString := fmt.Sprintf("%v", required)
  expectedString := fmt.Sprintf("%v", expected)
  if requiredString !=  expectedString {
    t.Errorf("Expected %v but got %v", expectedString, requiredString)
  }
}

func TestRequestAndParseDataComposition(t *testing.T) {
  _, codes := RequestLocationFromCoords(47.598755, -122.332764)
  r := newCompositionReport(t)
  result := r.RequestAndParseData(codes).(*CompositionConfigFormat)
  if result.Kind != "composition" {
    t.Errorf(`Kind should be "composition", but was %s.`, result.Kind)
  }
  if len(result.Parts) != 5 {
    t.Errorf("Parts should have 5 results, but instead has %d.", len(result.Parts))
  }
}