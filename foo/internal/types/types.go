package types

type Ad struct {
	ID       string `json:"id"`
	Customer string `json:"customer"`
	Info     string `json:"info"`
}

type FooRequest struct {
	Customer string `json:"customre,omitempty"`
}

type ZedRequest struct {
	Customer string `json:"customer,omitempty"`
}

type AdResponse struct {
	Ads   []Ad `json:"ads"`
	Count uint `json:"count"`
}
