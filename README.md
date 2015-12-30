# factual-go-driver

### Usage

```go
import (
  "net/url"
  
  "github.com/ainsleyc/factual"
  "github.com/bitly/go-simplejson"
)
  
// Initialization
key := "Factual API key"
secret := "Factual API secret"
client := factual.NewClient(key, secret)
    
// Get request
path := "/t/places-us"
params := url.Values{}
params.Set("q": "starbucks")
resp, err := client.Get(path, params)
    
// Parse response
respJson, err := simplejson.NewJson(resp)
data := respJson.Get("response").Get("data")
```
