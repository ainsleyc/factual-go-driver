// +build integration 

package factual_test

import (
  "os"
  "encoding/json"
  "testing"
  "net/url"

  "github.com/ainsleyc/factual"
)

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

func TestGet_ValidUrl_ShouldNotReturnError(t *testing.T) {
  path := "/t/place-categories"
  config, _:= getTestConfig()
  params := url.Values{}
  client := factual.NewClient(config.Key, config.Secret) 
  _, err := client.Get(path, params)
  if err != nil {
    t.Error("Get returned error for valid url, Factual API may be unavailable")
  }
}

