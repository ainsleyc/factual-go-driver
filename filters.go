package factual

import (
  "github.com/bitly/go-simplejson"
)

type LogicalOperator string

const (
  And LogicalOperator = "$and"
  Or LogicalOperator = "$or"
)

type ComparisonOperator string

const (
  Blank ComparisonOperator = "$blank" 
  Bw ComparisonOperator = "$bw" 
  Bwin ComparisonOperator = "$bwin" 
  Eq ComparisonOperator = "$eq" 
  Excludes ComparisonOperator = "$excludes" 
  ExcludesAny ComparisonOperator = "$excludes_any" 
  Gt ComparisonOperator = "$gt" 
  Gte ComparisonOperator = "$gte" 
  Includes ComparisonOperator = "$includes" 
  IncludesAny ComparisonOperator = "$includes_any" 
  Lt ComparisonOperator = "$lt" 
  Lte ComparisonOperator = "$lte" 
  Nbw ComparisonOperator = "$nbw" 
  Nbwin ComparisonOperator = "$bnwin" 
  Neq ComparisonOperator = "$neq" 
  Nin ComparisonOperator = "$nin" 
  Search ComparisonOperator = "$search" 
)

type FilterInterface interface {
  MarshalJSON() ([]byte, error)
  toJson() (*simplejson.Json, error)
}

type filter struct {
  Field string
  Op ComparisonOperator
  Vals interface{}
}

func NewFilter(field string, op ComparisonOperator, vals interface{}) *filter {
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
//   Op LogicalOperator 
//   Vals []FilterInterface
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
