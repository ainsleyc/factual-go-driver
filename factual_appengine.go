package factual

// +build appengine
import (
	"net/http"

	"golang.org/x/net/context"
	"google.golang.org/appengine"
	"google.golang.org/appengine/urlfetch"
)

func init() {
	// getClient retrieves an http.Client from the appengine
	// context
	getClient = func(ctx context.Context) *http.Client {
		return urlfetch.Client(ctx)
	}

}

// WithRequest configures an option to modify the internal
// http.Client of factual driver
func WithRequest(r *http.Request) Option {
	return func(c *Client) error {
		c.ctx = appengine.NewContext(r)
		return nil
	}
}
