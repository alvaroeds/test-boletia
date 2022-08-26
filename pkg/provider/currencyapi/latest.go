package currencyapi

import "time"

type Latest struct {
	Meta struct {
		LastUpdatedAt time.Time `json:"last_updated_at"`
	} `json:"meta"`
	Data map[string]Currency `json:"data"`
}

type Currency struct {
	Code  string  `json:"code"`
	Value float64 `json:"value"`
}

type QueryInfo struct {
	Method string
	Path   string
	Code   int
	Time   float64
}
