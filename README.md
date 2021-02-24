Web Lib v1.0.0
==============

Simple Go stdlib http.Client using defaults based on [Don’t use Go’s default HTTP client (in production) | by Nathan Smith | Medium][], and [How to Use the HTTP Client in GO To Enhance Performance · LoginRadius Engineering][]. You can also change the defaults.


Rational
--------

I often need a more sane default HTTP client that what is provided by Go. So here it is. Mostly for my own use but I do hope others find it useful as well.


Example
-------

```go
package main

import (
	"io/ioutil"
	"log"

	"github.com/runeimp/web"
)

func main() {
	web.Config.ClientTimeout = 30 // All Timeout values are in seconds

	httpClient := web.NewClient()
	resp, err := httpClient.Get("http://webtools.zone/phptool.php?phptool_form=get&phptool_act=http_headers")
	if err != nil {
		log.Fatalln(err)
	}

	// We Read the response body on the line below.
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatalln(err)
	}

	// Convert the body to type string
	sb := string(body)
	log.Println(sb)
}
```





[Don’t use Go’s default HTTP client (in production) | by Nathan Smith | Medium]: https://medium.com/@nate510/don-t-use-go-s-default-http-client-4804cb19f779
[How to Use the HTTP Client in GO To Enhance Performance · LoginRadius Engineering]: https://www.loginradius.com/blog/async/tune-the-go-http-client-for-high-performance/

