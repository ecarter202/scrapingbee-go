package scrapingbee

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"strconv"
	"strings"
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
	if params != nil && params.JavascriptRendering {
		values.Set("render_js", btos(params.JavascriptRendering))
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
	} else {
		values.Set("render_js", "false")
	}

	if params != nil && len(params.headers) > 0 {
		values.Set("forward_headers", "true")
	}
	if params != nil && len(params.cookies) > 0 {
		var cookieX []string
		for k, v := range params.cookies {
			cookieX = append(cookieX, k+"="+v)
		}
		values.Set("cookies", strings.Join(cookieX, ";"))
	}

	uri.RawQuery = values.Encode()

	req, err := http.NewRequest("GET", uri.String(), nil)
	if err != nil {
		return nil, err
	}
	if params != nil {
		for hk, hv := range params.headers {
			req.Header.Add(hk, hv)
		}
	}

	return c.Do(req)
}

// Usage gets the usage stats for your account.
func (c *Client) Usage() (stats *UsageStats, err error) {
	uri := fmt.Sprintf("%s%s?api_key=%s", baseURL, "usage", c.apikey)

	res, err := http.Get(uri)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	if err := json.NewDecoder(res.Body).Decode(&stats); err != nil {
		return nil, err
	}

	return stats, nil
}

func btos(b bool) (s string) {
	if b {
		return "true"
	}
	return "false"
}
