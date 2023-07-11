package types

type Ad struct {
	ID       int64  `json:"id"`
	Customer string `json:"customer"`
	Info     string `json:"info"`
}

type AdRequest struct {
	ID       int64  `json:id,omitempty`
	Customer string `json:"customer,omitempty"`
}

type AdResponse struct {
	Ads   []Ad `json:"ads"`
	Count int  `json:"count"`
}
