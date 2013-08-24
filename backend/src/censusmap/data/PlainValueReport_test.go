package data

import (
  "testing"
  "fmt"
  "sort"
)

const dummyConfig string = `{"kind": "plain_value", "vars": {"Total Population": "B01003_001", "Other": "B02001_007"}}`

func TestParseConfig(t *testing.T) {
  r := new(PlainValueReport)
  bytes := []byte(dummyConfig)
  r.ParseConfig(bytes)
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