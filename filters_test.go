// +build unit

package factual_test

import (
  "testing"

  "github.com/ainsleyc/factual"
)

func TestFilter_MarshalJson_Should(t *testing.T) {
  tests := []struct {
    field string
    op factual.ComparisonOperator
    vals interface{}
    expected string
  }{
    {"name", factual.Eq, "Factual", "{\"name\":{\"$eq\":\"Factual\"}}"},
  }

  for _, test := range tests {
    filter := factual.NewFilter(test.field, test.op, test.vals)
    bytes, _ := filter.MarshalJSON()
    if string(bytes) != test.expected {
      t.Error(string(bytes), "!=", test.expected)
    }
  }
}
