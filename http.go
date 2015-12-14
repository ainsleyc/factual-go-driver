package factual

import (
  "net/http"

  "github.com/asaskevich/govalidator"
)

const baseUrl = "http://api.factual.com"

func Get(path string) (string, error) {

  fullUrl := baseUrl + path 
  if !govalidator.IsURL(fullUrl) {
    return "", ErrInvalidUrl(fullUrl)
  }

  resp, err := http.Get(fullUrl)
  if err != nil {
    return "", err
  }
  defer resp.Body.Close()

  if resp.StatusCode != 200 {
    return "", ErrHttpResponse(fullUrl, resp.StatusCode, "BLAH")
  }

  return "BLAH", nil 
}

func Post(url string) (string, error) {
  return "POST", nil
}
