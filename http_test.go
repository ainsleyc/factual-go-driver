// +build unit

package factual_test

import (
  "net/url"
  "testing"

  "github.com/ainsleyc/factual"
)

func TestGet_InvalidUrl_ShouldReturnError(t *testing.T) {
  invalidPath := "http://blah.com/places"
  params := url.Values{}
  client := factual.NewClient("blah", "blah")
  _, err := client.Get(invalidPath, params)
  if err == nil {
    t.Error("Did not return error for invalid path")
  }
}
