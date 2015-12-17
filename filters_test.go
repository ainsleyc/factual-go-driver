// +build unit

package factual_test

import (
  "fmt"
  "testing"

  "github.com/ainsleyc/factual"
)

func TestFilter_ToJson_Should(t *testing.T) {
  filter := factual.Filter{"name", factual.Eq, "Factual"}
  bytes, _ := filter.ToJson().EncodePretty()
  fmt.Println(string(bytes))
}
