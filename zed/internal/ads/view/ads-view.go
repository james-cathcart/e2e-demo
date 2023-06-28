package view

import (
	"e2ezed/internal/ads"
	"e2ezed/internal/types"
	"encoding/json"
	"io"
	"log"
	"net/http"
)

type AdsView struct {
	dao ads.DAO
}

func New(dao ads.DAO) ads.API {
	return &AdsView{
		dao: dao,
	}
}

func (view *AdsView) GetAds(w http.ResponseWriter, r *http.Request) {

	log.Println(`GetAds invoked`)
	defer r.Body.Close()

	bodyBytes, err := io.ReadAll(r.Body)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	var adRequest types.AdRequest
	err = json.Unmarshal(bodyBytes, &adRequest)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	var ads []types.Ad

	if adRequest.Customer == `all` {
		ads, err = view.dao.GetAllAds()
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	} else {
		ads, err = view.dao.GetAdsByCustomer(adRequest.Customer)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	}

	adResponse := types.AdResponse{
		Ads:   ads,
		Count: len(ads),
	}

	jsonBytes, err := json.Marshal(adResponse)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set(`Content-Type`, `application/json`)
	w.WriteHeader(http.StatusOK)
	w.Write(jsonBytes)

}
