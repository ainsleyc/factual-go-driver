// +build unit

package factual_test

import (
  "testing"

  "github.com/ainsleyc/factual"
)

func TestRead_ReadPath_ShouldReturnCorrectPath(t *testing.T) {
  tests := []struct {
    opts factual.ReadOpts
    expected string
  }{
    {
      factual.ReadOpts{"places-us", nil},
      "/t/places-us",
    },
    {
      factual.ReadOpts{
        "places", 
        factual.NewFilter(
          "name", 
          factual.Eq, 
          "Factual", 
        ),
      },
      "/t/places?filters={\"name\":{\"$eq\":\"Factual\"}}",
    },
  }

  client := factual.NewClient("blah", "blah")
  for _, test := range tests {
    path, _ := client.ReadPath(test.opts)
    if path != test.expected {
      t.Error(path, "!=", test.expected)
    }
  }
}
