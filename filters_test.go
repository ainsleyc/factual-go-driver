// +build unit

package factual_test

import (
  "testing"

  "github.com/ainsleyc/factual"
)

func TestComparisonFilter_MarshalJson_ShouldReturnJsonString(t *testing.T) {
  tests := []struct {
    field string
    op factual.ComparisonOperator
    vals interface{}
    expected string
  }{
    {"name", factual.Eq, "Factual", "{\"name\":{\"$eq\":\"Factual\"}}"},
  }

  for _, test := range tests {
    filter := factual.NewComparisonFilter(test.field, test.op, test.vals)
    bytes, _ := filter.MarshalJSON()
    if string(bytes) != test.expected {
      t.Error(string(bytes), "!=", test.expected)
    }
  }
}

func TestLogicalFilter_MarshalJson_ShouldReturnJsonString(t *testing.T) {
  tests := []struct {
    op factual.LogicalOperator
    vals []factual.Filter
    expected string
  }{
    {
      factual.And, 
      []factual.Filter{
        factual.NewComparisonFilter(
          "name",
          factual.Eq,
          "starbucks",
        ),
        factual.NewComparisonFilter(
          "locality",
          factual.Eq,
          "los angeles",
        ),
      },
      "{\"$and\":[{\"name\":{\"$eq\":\"starbucks\"}},{\"locality\":{\"$eq\":\"los angeles\"}}]}",
    },
  }

  for _, test := range tests {
    filter := factual.NewLogicalFilter(test.op, test.vals)
    bytes, _ := filter.MarshalJSON()
    if string(bytes) != test.expected {
      t.Error(string(bytes), "!=", test.expected)
    }
  }
}
