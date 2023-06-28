package view

import (
	"e2ebar/internal/ads"
	"e2ebar/internal/types"
	"encoding/json"
	"io"
	"net/http"
)

type AdsView struct {
	svc ads.Service
}

func New(svc ads.Service) ads.API {
	return &AdsView{
		svc: svc,
	}
}

func (api *AdsView) Create(w http.ResponseWriter, r *http.Request) {

	defer r.Body.Close()
	bodyBytes, err := io.ReadAll(r.Body)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	var barRequest types.BarRequest
	err = json.Unmarshal(bodyBytes, &barRequest)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	api.svc.Create(barRequest.Ads...)

	w.WriteHeader(http.StatusCreated)
}
