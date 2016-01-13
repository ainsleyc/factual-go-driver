// +build integration

package factual_test

import (
	"encoding/json"
  "fmt"
	"net/url"
	"os"
	"testing"

	"github.com/ainsleyc/factual"
	"github.com/bitly/go-simplejson"
)

const testValidPath = "/t/place-categories"

var testEmptyParams = url.Values{}

type testConfig struct {
	Key    string
	Secret string
}

func getTestConfig() (conf testConfig, err error) {
	config := testConfig{}
	file, err := os.Open("conf.json")
	if err != nil {
		return config, err
	}

	decoder := json.NewDecoder(file)
	err = decoder.Decode(&config)
	if err != nil {
		return config, err
	}

	return config, nil
}

func testRead(t *testing.T, path string, params url.Values) {
	config, _ := getTestConfig()
	client := factual.NewClient(config.Key, config.Secret)

	resp, err := client.Get(path, params)
	if err != nil {
		t.Error("Get returned error for valid url, Factual API may be unavailable")
	}

	respJson, _ := simplejson.NewJson(resp)
	data := respJson.Get("response").Get("data")
	if len(data.MustArray()) <= 0 {
		t.Error("Valid Get query returned no results")
	} else {
    paramStr, _ := json.Marshal(params)
    fmt.Println("=== RESULTS:", len(data.MustArray()), "results for", path, string(paramStr))
  }
}

func testGeotag(t *testing.T, path string, params url.Values) {
	config, _ := getTestConfig()
	client := factual.NewClient(config.Key, config.Secret)

	resp, err := client.Get(path, params)
	if err != nil {
		t.Error("Get returned error for valid url, Factual API may be unavailable")
	}

	respJson, _ := simplejson.NewJson(resp)
	data := respJson.Get("response").Get("data")
	if (data.Get("country") == nil) {
		t.Error("Valid Get query returned no results")
	} else {
    paramStr, _ := json.Marshal(params)
    fmt.Println("=== RESULTS:", data.MustMap(), "for", path, string(paramStr))
  }
}

// Test existence of valid config.json file
func TestGet_ConfigFile_ShouldExistAndBeValid(t *testing.T) {
	config, err := getTestConfig()
	if err != nil {
		switch err.(type) {
		default:
			t.Error("conf.json has an unknown error")
		case *os.PathError:
			t.Error("conf.json does not exist")
		case *json.SyntaxError:
			t.Error("conf.json is not a valid json")
		}
	}
	if config.Key == "" {
		t.Error("conf.json is missing Key")
	}
	if config.Secret == "" {
		t.Error("conf.json is missing Secret")
	}
}

// /t/places-us?q=starbucks
func TestGet_ReadWithQuery_ShouldReturnResults(t *testing.T) {
	path := "/t/places-us"
	params := url.Values{}
	params.Set("q", "starbucks")

	testRead(t, path, params)
}

// /t/places-us?q=starbucks
func TestGet_ReadWithKey_ShouldReturnResults(t *testing.T) {
  config, _ := getTestConfig()
	path := "/t/places-us"
	params := url.Values{}
	params.Set("q", "starbucks")
	params.Set("KEY", config.Key)

	testRead(t, path, params)
}

// /t/places-us?filters={"name":{"$eq":"starbucks"}}
func TestGet_ReadWithSingleComparisonFilter_ShouldReturnResults(t *testing.T) {
	path := "/t/places-us"
	params := url.Values{}
	filters, _ := factual.NewComparisonFilter(
		"name",
		factual.Eq,
		"starbucks",
	).MarshalJSON()
	params.Set("filters", string(filters))

	testRead(t, path, params)
}

// /t/places-us?filters={"$and":[{"name":"starbucks"},{"locality":"new york"}]}
func TestGet_ReadWithLogicalFilter_ShouldReturnResults(t *testing.T) {
	path := "/t/places-us"
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

	testRead(t, path, params)
}

// /t/places-us?geo={"$circle":{"$center":[34.06021,-118.41828],"$meters":50}
func TestGet_ReadWithGeoCircle_ShouldReturnResults(t *testing.T) {
	path := "/t/places-us"
	params := url.Values{}
	geo, _ := factual.NewGeoCircle(
		float64(34.06021),
		float64(-118.41828),
		50,
	).MarshalJSON()
	params.Set("geo", string(geo))

	testRead(t, path, params)
}

// /t/places-us?geo="{"$within":{"$rect":[[34.0611,-118.42283],[34.05771,-118.41399]]}}
func TestGet_ReadWithGeoRect_ShouldReturnResults(t *testing.T) {
	path := "/t/places-us"
	params := url.Values{}
	geo, _ := factual.NewGeoRect(
		float64(34.06110),
		float64(-118.42283),
		float64(34.05771),
		float64(-118.41399),
	).MarshalJSON()
	params.Set("geo", string(geo))

	testRead(t, path, params)
}

// /geotag?latitude=37.782137&longitude=-122.405803&KEY=key
func TestGet_Geotag_ShouldReturnResults(t *testing.T) {
  config, _ := getTestConfig()
	path := "/geotag"
	params := url.Values{}
	params.Set("latitude", "37.782137")
	params.Set("longitude", "-122.405803")
	params.Set("KEY", config.Key)

	testGeotag(t, path, params)
}
