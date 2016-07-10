package factual

import (
	"encoding/json"
	"net/url"
	"os"
	"testing"

	"golang.org/x/net/context"
	// "github.com/ainsleyc/factual-go-driver/testdata"
)

var getTestClient func(t *testing.T) (context.Context, func())

func setupTest(t *testing.T, path string, params url.Values) ([]byte, error) {
	config, err := getTestConfig()
	if err != nil {
		t.Fatalf("failed to load test config: %s", err)
	}
	// TODO: make a mock API server that serves out of testdata
	// return testdata.Data[params["q"][0]], nil

	client := NewClient(config.Key, config.Secret)
	if getTestClient != nil {
		ctx, done := getTestClient(t)
		defer done()

		client.ctx = ctx
	}

	return client.Get(path, params)
}

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
