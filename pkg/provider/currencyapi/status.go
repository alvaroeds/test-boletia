package currencyapi

type Status struct {
	Quotas struct {
		Month struct {
			Total     int `json:"total"`
			Used      int `json:"used"`
			Remaining int `json:"remaining"`
		} `json:"month"`
		Grace struct {
			Total     int `json:"total"`
			Used      int `json:"used"`
			Remaining int `json:"remaining"`
		} `json:"grace"`
	} `json:"quotas"`
	Ratelimit int
}
