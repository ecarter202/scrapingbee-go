# scrapingbee-go
Taps into ScrapingBee API.

## Installing

Install the Go way:

```sh
go get -u github.com/ecarter202/scrapingbee-go
```

## Using

`````go
package main

import (
	"fmt"
  "io/ioutil"

	"github.com/ecarter202/scrapingbee-go"
)

func main() {
	client := scrapingbee.New("myapikey")
  
  params := &scrapingbee.ReqParams{
    JavascriptRendering:  true,
    PremiumProxies:       true,
    BlockAds:             true,
  }
  params.AddHeader("X-Custom", "myvalue")
  
  res, _ := client.Get("https://newegg.com", params)
  ...
}
`````
