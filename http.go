package factual

import (
  "io/ioutil"
  "net/http"
  // "net/url"
  // "strings"

  "github.com/asaskevich/govalidator"
  // "github.com/garyburd/go-oauth/oauth"
)

func (c Client) Get(path string) ([]byte, error) {

  fullUrl := c.BaseUri + path 
  if !govalidator.IsURL(fullUrl) {
    return nil, ErrInvalidUrl(fullUrl)
  }

  // form := url.Values{}
  // req, err := http.NewRequest("GET", fullUrl, strings.NewReader(form.Encode()))
  // if err != nil {
  //   return nil, err
  // }

  // resp, err := http.DefaultClient.Do(req)
  // if err != nil {
  //   return nil, err
  // }
  // defer resp.Body.Close()

  // resp, err := c.Oauth.Get(http.DefaultClient, &c.Creds, fullUrl, form)
  resp, err := c.Oauth.Get(http.DefaultClient, nil, fullUrl, nil)

  body, err := ioutil.ReadAll(resp.Body)
  if err != nil {
    return nil, ErrHttpBody(fullUrl) 
  }

  if resp.StatusCode != 200 {
    return nil, ErrHttpResponse(fullUrl, resp.StatusCode, body) 
  }

  return body, nil 
}

func Post(url string) (string, error) {
  return "POST", nil
}
