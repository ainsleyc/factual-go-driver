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
}

// type LogicalOperator string

// const (
// 	And LogicalOperator = "$and"
// 	Or  LogicalOperator = "$or"
// )

// type ComparisonOperator string

// const (
// 	Blank       ComparisonOperator = "$blank"
// 	Bw          ComparisonOperator = "$bw"
// 	Bwin        ComparisonOperator = "$bwin"
// 	Eq          ComparisonOperator = "$eq"
// 	Excludes    ComparisonOperator = "$excludes"
// 	ExcludesAny ComparisonOperator = "$excludes_any"
// 	Gt          ComparisonOperator = "$gt"
// 	Gte         ComparisonOperator = "$gte"
// 	Includes    ComparisonOperator = "$includes"
// 	IncludesAny ComparisonOperator = "$includes_any"
// 	Lt          ComparisonOperator = "$lt"
// 	Lte         ComparisonOperator = "$lte"
// 	Nbw         ComparisonOperator = "$nbw"
// 	Nbwin       ComparisonOperator = "$bnwin"
// 	Neq         ComparisonOperator = "$neq"
// 	Nin         ComparisonOperator = "$nin"
// 	Search      ComparisonOperator = "$search"
// )

// type Filter interface {
// 	MarshalJSON() ([]byte, error)
// 	toJson() *simplejson.Json
// }

// type ComparisonFilter struct {
// 	Field string
// 	Op    ComparisonOperator
// 	Vals  interface{}
// }

// func NewComparisonFilter(field string, op ComparisonOperator, vals interface{}) *ComparisonFilter {
// 	return &ComparisonFilter{field, op, vals}
// }

// func (f *ComparisonFilter) toJson() *simplejson.Json {
// 	opJson := simplejson.New()
// 	opJson.Set(string(f.Op), f.Vals)
// 	json := simplejson.New()
// 	json.Set(f.Field, opJson)
// 	return json
// }

// func (f *ComparisonFilter) MarshalJSON() ([]byte, error) {
// 	filter := f.toJson()
// 	bytes, err := filter.MarshalJSON()
// 	return bytes, err
// }

// type LogicalFilter struct {
// 	Op   LogicalOperator
// 	Vals []Filter
// }

// func NewLogicalFilter(op LogicalOperator, vals []Filter) *LogicalFilter {
// 	return &LogicalFilter{op, vals}
// }

// func (f *LogicalFilter) toJson() *simplejson.Json {
// 	json := simplejson.New()
// 	json.Set(string(f.Op), f.Vals)
// 	return json
// }

// func (f *LogicalFilter) MarshalJSON() ([]byte, error) {
// 	filter := f.toJson()
// 	bytes, err := filter.MarshalJSON()
// 	return bytes, err
// }
