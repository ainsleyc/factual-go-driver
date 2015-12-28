package factual

import (
	// "net/http"
	"github.com/garyburd/go-oauth/oauth"
)

type Client struct {
	Oauth   oauth.Client
	BaseUri string
}

const defaultUri = "http://api.v3.factual.com"

func NewClient(key string, secret string) Client {
	creds := oauth.Credentials{key, secret}
	oauthClient := oauth.Client{
		Credentials:     creds,
		SignatureMethod: oauth.HMACSHA1,
	}
	return Client{oauthClient, defaultUri}
}
