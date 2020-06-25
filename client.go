package scrapingbee

import (
	"net/http"
	"net/url"
	"strconv"
)

const (
	baseURL = "https://app.scrapingbee.com/api/v1/"
)

type Client struct {
	*http.Client
	apikey string
}

// New creates a new API client.
func New(apikey string) (c *Client) {
	return NewCustomClient(apikey, new(http.Client))
}

// NewCustomClient creates a new API client, based on the provided http.Client.
func NewCustomClient(apikey string, client *http.Client) (c *Client) {
	return &Client{
		apikey: apikey,
		Client: client,
	}
}

// Get will send a get request to the supplied target (url).
// Params are optional (defaults will be used if not supplied).
func (c *Client) Get(target string, params *ReqParams) (res *http.Response, err error) {
	uri, err := url.ParseRequestURI(baseURL)
	if err != nil {
		return nil, err
	}
	values := uri.Query()
	values.Set("api_key", c.apikey)
	values.Set("url", target)
	if params.JavascriptRendering {
		values.Set("render_us", btos(params.JavascriptRendering))
		if params.PremiumProxies {
			values.Set("premium_proxy", "true")
			if string(params.CountryCode) != "" {
				values.Set("country_code", string(params.CountryCode))
			}
		}
		if params.JavascriptSnippet != "" {
			values.Set("js_snippet", params.JavascriptSnippet)
		}
		if params.Wait.Milliseconds() > 0 {
			values.Set("wait", strconv.Itoa(int(params.Wait.Milliseconds())))
		}
		if params.WaitForCSSSel != "" {
			values.Set("wait_for", params.WaitForCSSSel)
		}
		if params.ReturnSource {
			values.Set("return_page_source", "true")
		}
		if params.BlockAds {
			values.Set("block_ads", "true")
		}
		if params.BlockAds {
			values.Set("block_resources", "true")
		}
	}

	uri.RawQuery = values.Encode()

	req, err := http.NewRequest("GET", uri.String(), nil)
	if err != nil {
		return nil, err
	}
	for hk, hv := range params.headers {
		req.Header.Add(hk, hv)
	}
	for ck, cv := range params.cookies {
		req.AddCookie(&http.Cookie{
			Name:  ck,
			Value: cv,
		})
	}

	return c.Do(req)
}

func btos(b bool) (s string) {
	if b {
		return "true"
	}
	return "false"
}
