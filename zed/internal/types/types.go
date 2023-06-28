package types

type Ad struct {
	ID       string `json:"id"`
	Customer string `json:"customer"`
	Info     string `json:"info"`
}

type AdRequest struct {
	Customer string `json:"customer,omitempty"`
}

type AdResponse struct {
	Ads   []Ad `json:"ads"`
	Count int  `json:"count"`
}
