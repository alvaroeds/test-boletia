package currency

import "time"

type GetAllCurrency struct {
	Code      string
	Value     float64
	CreatedAt time.Time
}

type InsertCurrency struct {
	Code  string
	Value float64
}

type CurrencyFilterRequest struct {
	FInit time.Time `json:"finit"`
	FEnd  time.Time `json:"fend"`
	Code  string
}
