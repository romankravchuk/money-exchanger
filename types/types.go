package types

type ConvertQuery struct {
	Amount float64 `json:"amount"`
	From   string  `json:"from"`
	To     string  `json:"to"`
}

type ConvertResponse struct {
	Query  ConvertQuery `json:"query"`
	Result float64      `json:"result"`
}
