// +build unit

package factual_test

import (
	"testing"

	"github.com/ainsleyc/factual"
)

func TestGeoCircle_MarshalJson_ShouldReturnJsonString(t *testing.T) {
	tests := []struct {
		lat      float64
		long     float64
		radius   int
		expected string
	}{
		{
			float64(34.06021),
			float64(-118.4184),
			int(5000),
			"{\"$circle\":{\"$center\":[34.06021,-118.4184],\"$meters\":5000}}",
		},
	}

	for _, test := range tests {
		geo := factual.NewGeoCircle(test.lat, test.long, test.radius)
		bytes, _ := geo.MarshalJSON()
		if string(bytes) != test.expected {
			t.Error(string(bytes), "!=", test.expected)
		}
	}
}

// func TestLogicalFilter_MarshalJson_ShouldReturnJsonString(t *testing.T) {
// 	tests := []struct {
// 		op       factual.LogicalOperator
// 		vals     []factual.Filter
// 		expected string
// 	}{
// 		{
// 			factual.And,
// 			[]factual.Filter{
// 				factual.NewComparisonFilter(
// 					"name",
// 					factual.Eq,
// 					"starbucks",
// 				),
// 				factual.NewComparisonFilter(
// 					"locality",
// 					factual.Eq,
// 					"los angeles",
// 				),
// 			},
// 			"{\"$and\":[{\"name\":{\"$eq\":\"starbucks\"}},{\"locality\":{\"$eq\":\"los angeles\"}}]}",
// 		},
// 	}

// 	for _, test := range tests {
// 		filter := factual.NewLogicalFilter(test.op, test.vals)
// 		bytes, _ := filter.MarshalJSON()
// 		if string(bytes) != test.expected {
// 			t.Error(string(bytes), "!=", test.expected)
// 		}
// 	}
// }
