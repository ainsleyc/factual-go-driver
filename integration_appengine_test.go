// +build appengine

package factual

import (
	"encoding/json"
	"net/url"
	"testing"

	"golang.org/x/net/context"

	"google.golang.org/appengine/aetest"
)

func init() {

	getTestClient = func(t *testing.T) (context.Context, func()) {
		ctx, done, err := aetest.NewContext()
		if err != nil {
			t.Fatal(err)
		}
		return ctx, done
	}
}

func TestGetClient_appengine(t *testing.T) {
	path := "/t/places-us"
	params := url.Values{}
	params.Set("q", "starbucks")

	bs, err := setupTest(t, path, params)
	if err != nil {
		t.Fatal(err)
	}

	// TODO: Client should return PlaceResponse
	var wrap PlaceWrap
	if err := json.Unmarshal(bs, &wrap); err != nil {
		t.Fatal(err)
	}

	resp := wrap.Response
	if resp.IncludedRows != len(resp.Data) {
		t.Errorf("got: %d wanted: %d", len(resp.Data),
			resp.IncludedRows)
	}
}
