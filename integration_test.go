// +build integration 

package factual

import (
  "testing"
  "net/url"
)


func TestGet_ValidUrl_ShouldNotReturnError(t *testing.T) {
  path := "/t/place-categories"
  params := url.Values{}
  client := NewClient("blah", "blah")
  _, err := client.Get(path, params)
  if err != nil {
    t.Error("Get returned error for valid url, Factual API may be unavailable")
  }
}

