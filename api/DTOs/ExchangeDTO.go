package dtos

type ExchangeRatesDTO struct {
	Success string  `json:"success"`
	Query   Query   `json:"query"`
	Info    Info    `json:"info"`
	Date    string  `json:"date"`
	Result  float64 `json:"result"`
}

type Query struct {
	From   string `json:"from"`
	To     string `json:"to"`
	Amount int64  `json:"amount"`
}

type Info struct {
	TimeStamp int64   `json:"timestamp"`
	Rate      float64 `json:"rate"`
}
