package factual

import (
  "github.com/bitly/go-simplejson"
)

type logicalOperator string

const (
  And logicalOperator = "$and"
  Or logicalOperator = "$or"
)

type comparisonOperator string

const (
  Blank comparisonOperator = "$blank" 
  Bw comparisonOperator = "$bw" 
  Bwin comparisonOperator = "$bwin" 
  Eq comparisonOperator = "$eq" 
  Excludes comparisonOperator = "$excludes" 
  ExcludesAny comparisonOperator = "$excludes_any" 
  Gt comparisonOperator = "$gt" 
  Gte comparisonOperator = "$gte" 
  Includes comparisonOperator = "$includes" 
  IncludesAny comparisonOperator = "$includes_any" 
  Lt comparisonOperator = "$lt" 
  Lte comparisonOperator = "$lte" 
  Nbw comparisonOperator = "$nbw" 
  Nbwin comparisonOperator = "$bnwin" 
  Neq comparisonOperator = "$neq" 
  Nin comparisonOperator = "$nin" 
  Search comparisonOperator = "$search" 
)

type filterInterface interface {
  MarshalJSON() ([]byte, error)
  toJson() (*simplejson.Json, error)
}

type filter struct {
  Field string
  Op comparisonOperator
  Vals interface{}
}

func NewFilter(field string, op comparisonOperator, vals interface{}) *filter {
  return &filter{field, op, vals}
}

func (f *filter) toJson() *simplejson.Json {
  opJson := simplejson.New()
  opJson.Set(string(f.Op), f.Vals)
  json := simplejson.New()
  json.Set(f.Field, opJson)
  return json
}

func (f *filter) MarshalJSON() ([]byte, error) {
  filter := f.toJson()
  bytes, err := filter.MarshalJSON()
  return bytes, err 
}

// type LogicalFilter struct {
//   Op logicalOperator 
//   Vals []filterInterface
// }

// func (f *LogicalFilter) toJson() *simplejson.Json {
//   json := simplejson.New()
//   json.Set(string(f.Op), "blah")
//   return json
// }

// func (f *Filter) MarshalJSON() ([]byte, error) {
//   filter := f.toJson()
//   bytes, err := filter.MarshalJSON()
//   return bytes, err 
// }
