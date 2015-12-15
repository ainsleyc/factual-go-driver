// +build unit

package factual

import (
  "net/url"
  "testing"
)

func TestGet_InvalidUrl_ShouldReturnError(t *testing.T) {
  invalidPath := "http://blah.com/places"
  params := url.Values{}
  client := NewClient("blah", "blah")
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

