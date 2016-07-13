package factual

import "testing"

func TestComparisonFilter_MarshalJson_ShouldReturnJsonString(t *testing.T) {
	tests := []struct {
		field    string
		op       ComparisonOperator
		vals     interface{}
		expected string
	}{
		{"name", Eq, "Factual", "{\"name\":{\"$eq\":\"Factual\"}}"},
	}

	for _, test := range tests {
		filter := NewComparisonFilter(test.field, test.op, test.vals)
		bytes, _ := filter.MarshalJSON()
		if string(bytes) != test.expected {
			t.Error(string(bytes), "!=", test.expected)
		}
	}
}

func TestLogicalFilter_MarshalJson_ShouldReturnJsonString(t *testing.T) {
	tests := []struct {
		op       LogicalOperator
		vals     []Filter
		expected string
	}{
		{
			And,
			[]Filter{
				NewComparisonFilter(
					"name",
					Eq,
					"starbucks",
				),
				NewComparisonFilter(
					"locality",
					Eq,
					"los angeles",
				),
			},
			"{\"$and\":[{\"name\":{\"$eq\":\"starbucks\"}},{\"locality\":{\"$eq\":\"los angeles\"}}]}",
		},
	}

	for _, test := range tests {
		filter := NewLogicalFilter(test.op, test.vals)
		bytes, _ := filter.MarshalJSON()
		if string(bytes) != test.expected {
			t.Error(string(bytes), "!=", test.expected)
		}
	}
}
