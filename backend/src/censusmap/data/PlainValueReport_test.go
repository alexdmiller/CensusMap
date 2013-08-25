package data

import (
  "testing"
  "fmt"
  "sort"
  "bytes"
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

func TestWriteFormattedReport(t *testing.T) {
  out := make([]byte, 0, 10)
  writer := bytes.NewBuffer(out)
  r := newPlainValueReport(t)
  r.setVariable("B01003_001E", "12345")
  r.setVariable("B02001_007E", "6798765")
  r.WriteFormattedReport(writer)
  expected := []byte(`{"kind":"plain_value","vars":{"Other":"6798765","Total Population":"12345"}}`)
  actual := writer.Bytes()
  if string(actual) != string(expected) {
    t.Error("Expected %s but got %s", expected, actual)
  }
}

func TestRequestDataAndWriteFormattedReport(t *testing.T) {
  _, codes := RequestLocationFromCoords(47.598755, -122.332764)
  r := newPlainValueReport(t)
  r.RequestData(codes)
  out := make([]byte, 0, 10)
  writer := bytes.NewBuffer(out)
  r.WriteFormattedReport(writer)
  actual := writer.Bytes()
  fmt.Printf("%s\n", actual)
}