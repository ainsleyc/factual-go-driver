# factual-go-driver

### Usage

```go
import (
  "net/url"
  
  "github.com/ainsleyc/factual"
  "github.com/bitly/go-simplejson"
)
  
// Initialization
key := "Factual API key"
secret := "Factual API secret"
client := factual.NewClient(key, secret)
    
// Get request
path := "/t/places-us"
params := url.Values{}
params.Set("q": "starbucks")
resp, err := client.Get(path, params)
    
// Parse response
respJson, err := simplejson.NewJson(resp)
data := respJson.Get("response").Get("data")
```

### Comparison Filters

```go
params := url.Values{}
filters, err := factual.NewComparisonFilter(
  "name",
  factual.Eq,
  "starbucks",
).MarshalJSON()
params.Set("filters", string(filters))
```

### Logical Filters

```go
params := url.Values{}
filter1 := factual.NewComparisonFilter(
  "name",
  factual.Eq,
  "starbucks",
)
filter2 := factual.NewComparisonFilter(
  "locality",
  factual.Eq,
  "new york",
)
andFilter, _ := factual.NewLogicalFilter(
  factual.And,
  []factual.Filter{filter1, filter2},
).MarshalJSON()
params.Set("filters", string(andFilter))
```
