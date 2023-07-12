package types

type Ad struct {
	ID       int64  `json:"id"`
	Customer string `json:"customer"`
	Info     string `json:"info"`
}

type FooRequest struct {
	Customer string `json:"customer,omitempty"`
}

type BarRequest struct {
	Ads []Ad
}

type AdResponse struct {
	Ads   []Ad `json:"ads"`
	Count int  `json:"count"`
}

type ETLRequest struct {
	Customer string `json:"customer,omitempty"`
	DryRun   bool   `json:"dryRun"`
}

type ETLResponse struct {
	Ads      []Ad `json:"ads"`
	Migrated int  `json:"migrated"`
}
