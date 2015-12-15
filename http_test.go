// +build unit

package factual_test

import (
  "net/url"
  "testing"
  // "encoding/json"

  "github.com/ainsleyc/factual"
)

func getTestConfig() (key string, secret string, err error) {
  return "", "", nil
}

func TestGet_ConfigFile_ShouldExist(t *testing.T) {

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

// func TestGet_HttpError_ShouldReturnError(t *testing.T) {
//   validPath := "/t/place-categories"
//   client := NewClient(fakeKey, fakeSecret)
//   _, err := client.Get(validPath)
//   if err == nil {
//     t.Error("Did not return error for http error code")
//   }
// }

