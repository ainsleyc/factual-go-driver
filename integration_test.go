// +build integration 

package factual_test

import (
  "os"
  "encoding/json"
  "testing"
  "net/url"

  "github.com/ainsleyc/factual"
  "github.com/bitly/go-simplejson"
)

const testValidPath = "/t/place-categories"
var testEmptyParams = url.Values{} 

type testConfig struct {
  Key string 
  Secret string
}

func getTestConfig() (conf testConfig, err error) {
  config := testConfig{}
  file, err := os.Open("conf.json")
  if err != nil {
    return config, err
  }

  decoder := json.NewDecoder(file)
  err = decoder.Decode(&config)
  if err != nil {
    return config, err
  }

  return config, nil
}

func testGet(t *testing.T, path string, params url.Values) {
  config, _:= getTestConfig()
  client := factual.NewClient(config.Key, config.Secret) 

  resp, err := client.Get(path, params)
  if err != nil {
    t.Error("Get returned error for valid url, Factual API may be unavailable")
  }

  json, _ := simplejson.NewJson(resp)
  data := json.Get("response").Get("data")
  if len(data.MustArray()) <= 0 {
    t.Error("Valid Get query returned no results")
  }
}

// Test existence of valid config.json file
func TestGet_ConfigFile_ShouldExistAndBeValid(t *testing.T) {
  config, err := getTestConfig()
  if err != nil {
    switch err.(type) {
    default:
      t.Error("conf.json has an unknown error")
    case *os.PathError:
      t.Error("conf.json does not exist")
    case *json.SyntaxError:
      t.Error("conf.json is not a valid json")
    }
  }
  if config.Key == "" {
    t.Error("conf.json is missing Key")
  }
  if config.Secret == "" {
    t.Error("conf.json is missing Secret")
  }
}

// /t/places-us?q=starbucks
func TestGet_ReadWithQuery_ShouldReturnResults(t *testing.T) {
  path := "/t/places-us" 
  params := url.Values{}
  params.Set("q", "starbucks")

  testGet(t, path, params)
}

// /t/places-us?filters={"name":{"$eq":"starbucks"}}
func TestGet_ReadWithSingleFilter_ShouldReturnResults(t *testing.T) {
  path := "/t/places-us" 
  params := url.Values{}
  filters, _ := factual.NewFilter(
    "name",
    factual.Eq,
    "starbucks",
  ).MarshalJSON()
  params.Set("filters", string(filters))

  testGet(t, path, params)
}

// /t/places-us?filters={"$and":[{"name":"starbucks"},{"locality":"los angeles"}]}
func TestGet_ReadWithLogicalFilter_ShouldReturnResults(t *testing.T) {
  path := "/t/places-us" 
  params := url.Values{}
  filter1, _ := factual.NewFilter(
    "name",
    factual.Eq,
    "starbucks",
  )
  filter2, _ := factual.NewFilter(
    "locality",
    factual.Eq,
    "los angeles",
  )
  andFilter, _ := factual.NewLogicalFilter(
    factual.And,
    []factual.Filter{filter1, filter2},
  ).MarshalJSON()
  params.Set("filters", string(andFilter))

  testGet(t, path, params)
}
