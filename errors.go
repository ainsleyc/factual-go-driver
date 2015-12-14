package factual

import (
  "fmt"
  "strconv"
  "strings"
  "errors"
)

func ErrInvalidUrl(url string) error {
  url += ": invalid url"
  return errors.New(url)
}

func ErrHttpResponse(url string, code int, message string) error {
  m := []string{url, " (", strconv.Itoa(code), "): ", message}
  fmt.Println(strings.Join(m, ""))
  return errors.New(strings.Join(m, ""))
}

