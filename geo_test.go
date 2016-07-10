package factual

import "testing"

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
		geo := NewGeoCircle(test.lat, test.long, test.radius)
		bytes, _ := geo.MarshalJSON()
		if string(bytes) != test.expected {
			t.Error(string(bytes), "!=", test.expected)
		}
	}
}

func TestGeoRect_MarshalJson_ShouldReturnJsonString(t *testing.T) {
	tests := []struct {
		tlLat    float64
		tlLong   float64
		brLat    float64
		brLong   float64
		expected string
	}{
		{
			float64(34.06110),
			float64(-118.42283),
			float64(34.05771),
			float64(-118.41399),
			"{\"$within\":{\"$rect\":[[34.0611,-118.42283],[34.05771,-118.41399]]}}",
		},
	}

	for _, test := range tests {
		geo := NewGeoRect(test.tlLat, test.tlLong, test.brLat, test.brLong)
		bytes, _ := geo.MarshalJSON()
		if string(bytes) != test.expected {
			t.Error(string(bytes), "!=", test.expected)
		}
	}
}
