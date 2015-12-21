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

func checkGetErr(t *testing.T, err error) {
  if err != nil {
    t.Error("Get returned error for valid url, Factual API may be unavailable")
  }
}

func checkRespHasData(t *testing.T, resp []byte) {
  json, _ := simplejson.NewJson(resp)
  data := json.Get("response").Get("data")
  if len(data.MustArray()) <= 0 {
    t.Error("Valid Get query returned no results")
  }
}

func TestGet_ConfigFile_ShouldExist(t *testing.T) {
  _, err := getTestConfig()
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
}

func TestGet_ConfigFile_ShouldHaveRequiredFields(t *testing.T) {
  config, _:= getTestConfig()
  if config.Key == "" {
    t.Error("conf.json is missing Key")
  }
  if config.Secret == "" {
    t.Error("conf.json is missing Secret")
  }
}

func TestGet_InvalidCredentials_ShouldReturnError(t *testing.T) {
  client := factual.NewClient("blah", "blah")
  _, err := client.Get(testValidPath, testEmptyParams)
  if err == nil {
    t.Error("Did not return error for invalid credentials")
  }
}

func TestGet_ValidUrl_ShouldNotReturnError(t *testing.T) {
  config, _:= getTestConfig()
  client := factual.NewClient(config.Key, config.Secret) 
  _, err := client.Get(testValidPath, testEmptyParams)
  checkGetErr(t, err)
}

func TestGet_ReadWithQuery_ShouldReturnResults(t *testing.T) {
  config, _:= getTestConfig()
  client := factual.NewClient(config.Key, config.Secret) 

  path := "/t/places-us" 
  params := url.Values{}
  params.Set("q", "starbucks")

  resp, err := client.Get(path, params)
  checkGetErr(t, err)
  checkRespHasData(t, resp)
}

func TestGet_ReadWithSingleFilter_ShouldReturnResults(t *testing.T) {
  config, _:= getTestConfig()
  client := factual.NewClient(config.Key, config.Secret) 

  path := "/t/places-us" 
  params := url.Values{}
  filters, _ := factual.NewFilter(
    "name",
    factual.Eq,
    "starbucks",
  ).MarshalJSON()
  params.Set("filters", string(filters))

  resp, err := client.Get(path, params)
  checkGetErr(t, err)
  checkRespHasData(t, resp)
}

