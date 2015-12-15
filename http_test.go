// +build unit

package factual_test

import (
  "fmt"
  "net/url"
  "testing"
  "os"
  "encoding/json"

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
  fmt.Println(config)
}

func TestGet_InvalidUrl_ShouldReturnError(t *testing.T) {
  invalidPath := "http://blah.com/places"
  params := url.Values{}
  client := factual.NewClient("blah", "blah")
  _, err := client.Get(invalidPath, params)
  if err == nil {
    t.Error("Did not return error for invalid path")
  }
}
