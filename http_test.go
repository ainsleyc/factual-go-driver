package factual

import (
  "testing"
)

const invalidPath = "http://blah.com/places"
const validPath = "/t/place-categories"

const fakeKey = "blah"
const fakeSecret = "blah"

func TestGet_InvalidUrl_ShouldReturnError(t *testing.T) {
  client := NewClient(fakeKey, fakeSecret)
  _, err := client.Get(invalidPath)
  if err == nil {
    t.Error("Did not return error for invalid path")
  }
}

func TestGet_HttpError_ShouldReturnError(t *testing.T) {
  client := NewClient(fakeKey, fakeSecret)
  _, err := client.Get(validPath)
  if err == nil {
    t.Error("Did not return error for http error code")
  }
}

// func TestGet_ValidUrl_ShouldNotReturnError(t *testing.T) {
//   _, err := Get("/t/place-categories")
//   if err != nil {
//     t.Error("Get returned error for valid url")
//   }
// }

