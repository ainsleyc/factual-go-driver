# factual-go-driver

### Get using OAuth

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

### Get using "KEY" parameter

```go
path := "/geotag"
params := url.Values{}
params.Set("latitude", "37.782137")
params.Set("longitude", "-122.405803")

// Setting "KEY" parameter converts request to non-oauth
params.Set("KEY", "Factual API key")

resp, err := client.Get(path, params)
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
andFilter, err := factual.NewLogicalFilter(
  factual.And,
  []factual.Filter{filter1, filter2},
).MarshalJSON()
params.Set("filters", string(andFilter))
```

### Geo Circle Filters

```go
params := url.Values{}
geo, err := factual.NewGeoCircle(
  float64(34.06021),   // Latitude
  float64(-118.41828), // Longitude
  50,                  // Radius (m)
).MarshalJSON()
params.Set("geo", string(geo))
```

### Geo Rectangle Filters

```go
params := url.Values{}
geo, err := factual.NewGeoRect(
  float64(34.06110),   // Upper left corner latitude
  float64(-118.42283), // Upper left corner longitude
  float64(34.05771),   // Bottom right corner latitude
  float64(-118.41399), // Bottom right corner longitude
).MarshalJSON()
params.Set("geo", string(geo))
```

```go
package main

import (
	"encoding/json"
	"log"
	"net/http"
	"net/url"
)

func main() {
	http.HandleFunc("/auto", auto)
	log.Fatal(http.ListenAndServe(":12345", nil))
}

func auto(w http.ResponseWriter, r *http.Request) {

	path := "/t/places-us"
	params := url.Values{}
	q := "ba"

	cmpFilter := driver.NewComparisonFilter(
		"name",
		driver.Bw,
		q,
	)

	bs, err := json.Marshal(cmpFilter)
	if err != nil {
		return httputils.NewServerErr(err)
	}

	params.Set("filters", string(bs))

	wc := driver.WithRequest(r)
	client := driver.NewClient(key, secret, wc)

	bs, err = client.Get(path, params)
	if err != nil {
		return httputils.NewServerErr(err)
	}

	w.Write(bs)
	w.WriteHeader(200)
}

```