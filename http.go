package factual

import (
	"errors"
	"io/ioutil"
	"net/http"
	"net/url"

	"github.com/asaskevich/govalidator"
)

func (c Client) Get(path string, params url.Values) ([]byte, error) {

	if params.Get("KEY") == "" {
		body, err := c.getOauth(path, params)
		return body, err
	}

	body, err := c.getWithKey(path, params)
	return body, err
}

func (c Client) getOauth(path string, params url.Values) ([]byte, error) {
	fullUrl := c.BaseUri + path
	if !govalidator.IsURL(fullUrl) {
		return nil, ErrInvalidUrl(fullUrl)
	}

	httpClient := http.DefaultClient
	if getClient != nil {
		if c.ctx == nil {
			return nil, errors.New("Appengine requires a context to be passed")
		}
		httpClient = getClient(c.ctx)
	}

	resp, err := c.Oauth.Get(httpClient, nil, fullUrl, params)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, ErrHttpBody(fullUrl)
	}

	if resp.StatusCode != 200 {
		return nil, ErrHttpResponse(fullUrl, resp.StatusCode, body)
	}

	return body, nil
}

func (c Client) getWithKey(path string, params url.Values) ([]byte, error) {
	fullUrl := c.BaseUri + path + "?" + params.Encode()

	resp, err := http.Get(fullUrl)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, ErrHttpBody(fullUrl)
	}

	if resp.StatusCode != 200 {
		return nil, ErrHttpResponse(fullUrl, resp.StatusCode, body)
	}

	return body, nil
}
