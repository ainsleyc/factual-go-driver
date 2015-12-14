package factual

import (
  "errors"
)

func ErrInvalidUrl(url string) error {
  url += " is an invalid url"
  return errors.New(url)
}

