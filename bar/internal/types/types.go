package types

type Ad struct {
	ID       string `json:"id"`
	Customer string `json:"customer"`
	Info     string `json:"info"`
}

type BarRequest struct {
	Ads []Ad
}
