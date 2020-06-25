package scrapingbee

import "time"

// ReqParams holds the parameters used in a request.
type ReqParams struct {
	headers map[string]string
	cookies map[string]string

	// JavascriptRendering fetches the URL through a real web browser.
	// Useful for website with lots of JavaScript.
	// Costs 5 credits
	// Learn more at: https://www.scrapingbee.com/documentation/#javascript-rendering
	JavascriptRendering bool

	// PremiumProxies utilizes residential proxies.
	// Only available when JS rendering is true.
	// Cost 100 credits.
	// Learn more at: https://www.scrapingbee.com/documentation/#premium-proxy
	PremiumProxies bool

	// CountryCode sets the country the request will originate from.
	// Currently only available with Premium Proxy.
	// Can work with any ISO-3166 country code.
	CountryCode countryCode

	// JavascriptSnippet is javascript code executed once the DOM is loaded.
	// Must be used with JavaScript rendering.
	// Learn more at: https://www.scrapingbee.com/documentation/#javascript-execution
	JavascriptSnippet string

	// Wait sets the amount of ms our server will wait for before returning your the result.
	// Useful for heavy websites.
	// Only works with JavaScript rendering.
	// Learn more at: https://www.scrapingbee.com/documentation/#wait
	Wait time.Duration

	// WaitForCSSSel is a CSS selector of the element our server will wait for before returning your the result.
	// Useful for heavy websites.
	// Only works with JavaScript rendering.
	// Learn more at: https://www.scrapingbee.com/documentation/#wait-for
	WaitForCSSSel string

	// ReturnSource will return set if you want to have the HTML returned by the server and unaltered by the browser.
	// Only available with JavaScript rendering.
	// Learn more at: https://www.scrapingbee.com/documentation/#page-source
	ReturnSource bool

	// BlockAds if you want to block ads on the page.
	// Only available with JavaScript rendering.
	// Learn more at: https://www.scrapingbee.com/documentation/#block-ads
	BlockAds bool

	// BlockResources if you want to block images and CSS on the page.
	// Only available with JavaScript rendering.
	// Learn more at: https://www.scrapingbee.com/documentation/#block-resources
	BlockResources bool
}

// AddHeader will add a header to the request parameters.
// Headers are forwarded to the website you want to scrape.
// Learn more at: https://www.scrapingbee.com/documentation/#header-forwarding
func (rp *ReqParams) AddHeader(key, value string) {
	rp.headers[key] = value
}

// AddCookie will add a cookie to the request parameters.
// Cookies are forwarded to the website you want to scrape.
// Learn more at: https://www.scrapingbee.com/documentation/#custom-cookies
func (rp *ReqParams) AddCookie(key, value string) {
	rp.cookies[key] = value
}
