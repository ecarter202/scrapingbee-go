package scrapingbee

// UsageStats are the statics for an account's usage.
type UsageStats struct {
	MaxAPICredit       int64 `json:"max_api_credit"`
	UsedAPICredit      int64 `json:"used_api_credit"`
	MaxConcurrency     int64 `json:"max_concurrency"`
	CurrentConcurrency int64 `json:"current_concurrency"`
}
