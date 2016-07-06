package factual

import (
	"net/http"

	"github.com/garyburd/go-oauth/oauth"
	"golang.org/x/net/context"
)

type Client struct {
	Oauth   oauth.Client
	BaseUri string
	ctx     context.Context
}

const defaultUri = "https://api.factual.com"

// NewClient accepts a client ID and client secret and returns
// a factual client. It accepts variable number of functional
// options for configuring the client.
func NewClient(key string, secret string, opts ...Option) Client {
	creds := oauth.Credentials{key, secret}
	oauthClient := oauth.Client{
		Credentials:     creds,
		SignatureMethod: oauth.HMACSHA1,
	}
	c := &Client{
		Oauth:   oauthClient,
		BaseUri: defaultUri,
	}
	for _, opt := range opts {
		opt(c)
	}
	return *c
}

// Option are functional options for configuring a Client
type Option func(c *Client) error

var getClient func(ctx context.Context) *http.Client
