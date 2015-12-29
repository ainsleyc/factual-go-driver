package factual

import (
	"github.com/bitly/go-simplejson"
)

type GeoOperator string

const (
	Circle GeoOperator = "$circle"
	Rect   GeoOperator = "$rect"
	Center GeoOperator = "$center"
	Within GeoOperator = "$within"
	Meters GeoOperator = "$meters"
)

type Geo interface {
	MarshalJSON() ([]byte, error)
	toJson() *simplejson.Json
}

type GeoCircle struct {
	lat    float64
	long   float64
	radius int
}

func NewGeoCircle(lat float64, long float64, radius int) *GeoCircle {
	return &GeoCircle{lat, long, radius}
}

func (g *GeoCircle) toJson() *simplejson.Json {
	opJson := simplejson.New()
	opJson.Set(string(Center), []float64{g.lat, g.long})
	opJson.Set(string(Meters), g.radius)
	json := simplejson.New()
	json.Set(string(Circle), opJson)
	return json
}

func (g *GeoCircle) MarshalJSON() ([]byte, error) {
	filter := g.toJson()
	bytes, err := filter.MarshalJSON()
	return bytes, err
}

type GeoRect struct {
	TlLat  float64
	TlLong float64
	BrLat  float64
	BrLong float64
}

func NewGeoRect(tlLat float64, tlLong float64,
	brLat float64, brLong float64) *GeoRect {
	return &GeoRect{tlLat, tlLong, brLat, brLong}
}

func (g *GeoRect) toJson() *simplejson.Json {
	coordinates := [][]float64{
		[]float64{g.TlLat, g.TlLong},
		[]float64{g.BrLat, g.BrLong},
	}
	opJson := simplejson.New()
	opJson.Set(string(Rect), coordinates)
	json := simplejson.New()
	json.Set(string(Within), opJson)
	return json
}

func (g *GeoRect) MarshalJSON() ([]byte, error) {
	filter := g.toJson()
	bytes, err := filter.MarshalJSON()
	return bytes, err
}
