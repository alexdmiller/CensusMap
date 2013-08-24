package data

import ("testing")

const dummyConfig string = `{"kind": "plain_value", "vars": {"Total Population": "B01003_001", "Other": "B02001_007"}}`

func TestParseConfig(t *testing.T) {
  r := new(PlainValueReport)
  bytes := []byte(dummyConfig)
  r.ParseConfig(bytes)
}