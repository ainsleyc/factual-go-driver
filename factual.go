package factual

type Client struct {
  key string
  secret string
  baseUri string
}

const baseUrl = "http://api.factual.com"

func NewClient(key string, secret string) Client {
  return Client{key, secret, baseUrl} 
}

