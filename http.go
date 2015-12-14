package factual

import (
  "io/ioutil"
  "net/http"

  "github.com/asaskevich/govalidator"
  // "github.com/garyburd/go-oauth/oauth"
)

// const baseUrl = "http://api.factual.com"

func (c Client) Get(path string) ([]byte, error) {

  fullUrl := baseUrl + path 
  if !govalidator.IsURL(fullUrl) {
    return nil, ErrInvalidUrl(fullUrl)
  }

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

func Post(url string) (string, error) {
  return "POST", nil
}
