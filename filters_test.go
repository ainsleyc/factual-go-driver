// +build unit

package factual_test

import (
  "fmt"
  "testing"

  "github.com/ainsleyc/factual"
)

func TestFilter_MarshalJson_Should(t *testing.T) {
  filter := factual.NewFilter("name", factual.Eq, "Factual")
  bytes, _ := filter.MarshalJSON()
  fmt.Println(string(bytes))
}
