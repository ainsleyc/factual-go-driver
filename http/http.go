package factual

import (
  "fmt"
  "net/http"

  "github.com/asaskevich/govalidator"
)

const baseUrl = "http://api.factual.com"

func Get(path string) (string, error) {
  
  fullUrl := baseUrl + path 
  if !govalidator.IsURL(fullUrl) {
    return "", ErrInvalidUrl 
  }

  resp, err := http.Get(fullUrl)
  if err != nil {
    return "", err
  }
  defer resp.Body.Close()

  fmt.Println("RESP:", resp.Body)

  return "BLAH", nil 
}

func Post(url string) (string, error) {
  return "POST", nil
}
