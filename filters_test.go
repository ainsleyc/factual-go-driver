// +build unit

package factual_test

import (
  "fmt"
  "testing"

  "github.com/ainsleyc/factual"
)

func TestGet_Json_Should(t *testing.T) {
  fmt.Println("HERE")
  // invalidPath := "http://blah.com/places"
  // params := url.Values{}
  // client := factual.NewClient("blah", "blah")
  // _, err := client.Get(invalidPath, params)
  // if err == nil {
  //   t.Error("Did not return error for invalid path")
  // }
  json := factual.JsonTest()
  bytes, _ := json.EncodePretty()
  fmt.Println(string(bytes))
  fmt.Println(factual.And)
}
