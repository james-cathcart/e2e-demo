package view

import (
	"e2efoo/internal/ads"
	"e2efoo/internal/types"
	"encoding/json"
	"io"
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

func (api *HttpView) GetAdsByCustomer(w http.ResponseWriter, r *http.Request) {

	log.Println(`GetAdsByCustomer endpoint called`)

	defer r.Body.Close()
	bodyBytes, err := io.ReadAll(r.Body)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	var fooRequest types.FooRequest
	err = json.Unmarshal(bodyBytes, &fooRequest)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	adsRecords, err := api.svc.GetAdsByCustomer(fooRequest)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	jsonBytes, err := json.Marshal(adsRecords)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set(`Content-Type`, `application/json`)
	w.WriteHeader(http.StatusOK)
	w.Write(jsonBytes)

}
