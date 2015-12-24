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

type filterInterface interface {
  MarshalJSON() ([]byte, error)
  toJson() (*simplejson.Json, error)
}

type Filter struct {
  Field string
  Op ComparisonOperator
  Vals interface{}
}

func NewFilter(field string, op ComparisonOperator, vals interface{}) *Filter {
  return &Filter{field, op, vals}
}

func (f *Filter) toJson() *simplejson.Json {
  opJson := simplejson.New()
  opJson.Set(string(f.Op), f.Vals)
  json := simplejson.New()
  json.Set(f.Field, opJson)
  return json
}

func (f *Filter) MarshalJSON() ([]byte, error) {
  filter := f.toJson()
  bytes, err := filter.MarshalJSON()
  return bytes, err 
}

type LogicalFilter struct {
  Op LogicalOperator 
  Vals *[]filterInterface
}

func NewLogicalFilter(op LogicalOperator, vals *[]filterInterface) *LogicalFilter {
  return &LogicalFilter{op, vals}
}

func (f *LogicalFilter) toJson() *simplejson.Json {
  // json := simplejson.New()
  // json.Set(string(f.Op), "blah")
  // return json
  return nil
}

func (f *LogicalFilter) MarshalJSON() ([]byte, error) {
  // filter := f.toJson()
  // bytes, err := filter.MarshalJSON()
  // return bytes, err 
  return nil, nil
}
