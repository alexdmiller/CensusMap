package data

import (
  "testing"
  "fmt"
  "sort"
  "bytes"
  "log"
)

const dummyConfig string = `{"kind": "plain_value", "vars": {"Total Population": "B01003_001", "Other": "B02001_007"}}`

func newPlainValueReport() *PlainValueReport {
  r := new(PlainValueReport)
  bytes := []byte(dummyConfig)
  r.ParseConfig(bytes)
  return r
}

func TestParseConfig(t *testing.T) {
  r := newPlainValueReport()
  required := r.GetRequiredVariables()
  expected := []string{"B01003_001", "B02001_007"}
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
  r := newPlainValueReport()
  r.SetVariable("B01003_001", "12345")
  r.SetVariable("B02001_007", "6798765")
  r.WriteFormattedReport(writer)
  expected := []byte(`{"kind":"plain_value","vars":{"Other":"6798765","Total Population":"12345"}}`)
  actual := writer.Bytes()
  if string(actual) != string(expected) {
    t.Error("Expected %s but got %s", expected, actual)
  }
}