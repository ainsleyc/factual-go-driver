package factual

import (
  "fmt"
  "testing"
)

const validPath = "/t/place-categories"
const invalidPath = "http://blah.com/places"

func TestGet_InvalidUrl_ShouldReturnError(t *testing.T) {
  _, err := Get(invalidPath)
  if err == nil {
    t.Error("Did not return error for invalid path")
  }
}

// func TestGet_ValidUrl(t *testing.T) {
//   _, err := Get("/t/place-categories")
//   if err != nil {
//     t.Error("Get returned error for valid Url")
//   }
// }

// func TestPost(t *testing.T) {
//   r, err := Post("BLAH")
// }
