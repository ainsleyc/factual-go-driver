// +build integration 

package factual

import (
  "testing"
)


func TestGet_ValidUrl_ShouldNotReturnError(t *testing.T) {
  var path = "/t/place-categories"
  client := NewClient("blah", "blah")
  _, err := client.Get(path)
  if err != nil {
    t.Error("Get returned error for valid url, Factual API may be unavailable")
  }
}

