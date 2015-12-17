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
  ToJson() (*simplejson.Json, error)
  MarshalJSON() ([]byte, error)
}

type Filter struct {
  Field string
  Op ComparisonOperator
  Vals interface{}
}

func (f *Filter) ToJson() *simplejson.Json {
  opJson := simplejson.New()
  opJson.Set(string(f.Op), f.Vals)
  json := simplejson.New()
  json.Set(f.Field, opJson)
  return json
}

type LogicalFilter struct {
  Op LogicalOperator 
  Vals []FilterInterface
}

