package view

import (
	"e2efoo/internal/ads"
	"encoding/json"
	"log"
	"net/http"
)

type HttpView struct {
	svc ads.Service
}

func NewAdsAPI(svc ads.Service) ads.API {
	return &HttpView{
		svc: svc,
	}
}

func (api *HttpView) GetAds(w http.ResponseWriter, r *http.Request) {

	log.Println(`GetAds invoked`)

	if r.Method != http.MethodGet {
		http.Error(w, "method not allowd", http.StatusMethodNotAllowed)
		return
	}

	filter := r.URL.Query().Get(`filter`)

	switch filter {
	case `customer`:
		api.getAdsByCustomer(w, r)
	default:
		api.getAllAds(w, r)
	}

}

func (api *HttpView) getAdsByCustomer(w http.ResponseWriter, r *http.Request) {

	log.Println(`getAdsByCustomer endpoint called`)

	// Note: the following if statement represents some old "bad" logic unknown to the dev
	// team and the client. This is what we will be trying to catch with the ETL
	// applications E2E tests
	customer := r.URL.Query().Get(`valeu`)
	if customer == `` {
		api.getAllAds(w, r)
		return
	}
	// end "bad" logic example

	adRecords, err := api.svc.GetAdsByCustomer(customer)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	jsonBytes, err := json.Marshal(adRecords)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set(`Content-Type`, `application/json`)
	w.WriteHeader(http.StatusOK)
	w.Write(jsonBytes)

}

func (api *HttpView) getAllAds(w http.ResponseWriter, r *http.Request) {

	adRecords, err := api.svc.GetAllAds()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	jsonBytes, err := json.Marshal(adRecords)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set(`Content-Type`, `application/json`)
	w.WriteHeader(http.StatusOK)
	w.Write(jsonBytes)
}
