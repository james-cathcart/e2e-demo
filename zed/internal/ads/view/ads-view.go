package view

import (
	"e2ezed/internal/ads"
	"e2ezed/internal/types"
	"encoding/json"
	"io"
	"log"
	"net/http"
	"strconv"
)

type AdsView struct {
	dao ads.DAO
}

func New(dao ads.DAO) ads.API {
	return &AdsView{
		dao: dao,
	}
}

func (view *AdsView) ServeHTTP(w http.ResponseWriter, r *http.Request) {

	switch r.Method {
	case http.MethodGet:
		view.GetAds(w, r)
	case http.MethodPost:
		view.CreateAd(w, r)
	case http.MethodPut:
		view.UpdateAd(w, r)
	case http.MethodDelete:
		view.DeleteAdByID(w, r)
	default:
		http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
		return
	}

}

func (view *AdsView) GetAds(w http.ResponseWriter, r *http.Request) {

	log.Println(`GetAds invoked`)

	filter := r.URL.Query().Get(`filter`)

	switch filter {
	case `id`:
		// ID lookup
		view.getAdsByID(w, r)
	case `customer`:
		// customer lookup
		view.getAdsByCustomer(w, r)
	default:
		// all ads
		view.getAllAds(w, r)
	}

}

func (view *AdsView) getAdsByID(w http.ResponseWriter, r *http.Request) {

	log.Println(`getAdsByID invoked`)

	idStr := r.URL.Query().Get(`value`)
	id32, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	id64 := int64(id32) // convert to long

	adSlice, err := view.dao.GetAdByID(id64)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	var adResponse types.AdResponse

	if adSlice == nil && err == nil {
		adResponse = types.AdResponse{
			Ads:   []types.Ad{},
			Count: 0,
		}
		w.WriteHeader(http.StatusNoContent)
	} else {
		adResponse = types.AdResponse{
			Ads:   adSlice,
			Count: len(adSlice),
		}
		w.WriteHeader(http.StatusOK)
	}

	jsonBytes, err := json.Marshal(adResponse)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set(`Content-Type`, `application/json`)
	w.Write(jsonBytes)

}

func (view *AdsView) getAdsByCustomer(w http.ResponseWriter, r *http.Request) {

	log.Println(`getAdsByCustomer invoked`)

	customer := r.URL.Query().Get(`value`)
	adSlice, err := view.dao.GetAdsByCustomer(customer)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	adResp := types.AdResponse{
		Ads:   adSlice,
		Count: len(adSlice),
	}

	jsonBytes, err := json.Marshal(adResp)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set(`Content-Type`, `application/json`)
	w.WriteHeader(http.StatusOK)
	w.Write(jsonBytes)

}

func (view *AdsView) getAllAds(w http.ResponseWriter, r *http.Request) {

	log.Println(`getAllAds invoked`)

	adSlice, err := view.dao.GetAllAds()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	adResp := types.AdResponse{
		Ads:   adSlice,
		Count: len(adSlice),
	}

	jsonBytes, err := json.Marshal(adResp)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Header().Set(`Content-Type`, `application/json`)
	w.Write(jsonBytes)

}

func (view *AdsView) CreateAd(w http.ResponseWriter, r *http.Request) {

	log.Println(`create ad invoked`)

	defer r.Body.Close()

	bodyBytes, err := io.ReadAll(r.Body)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	var newAd types.Ad
	err = json.Unmarshal(bodyBytes, &newAd)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	id, err := view.dao.CreateAd(newAd)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	response := make(map[string]int64)
	response[`id`] = id

	jsonBytes, err := json.Marshal(response)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	w.Header().Set(`Content-Type`, `application/json`)
	w.Write(jsonBytes)

}

func (view *AdsView) UpdateAd(w http.ResponseWriter, r *http.Request) {

	log.Println(`update ad invoked`)

	defer r.Body.Close()

	bodyBytes, err := io.ReadAll(r.Body)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	var updateReq types.Ad
	err = json.Unmarshal(bodyBytes, &updateReq)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	err = view.dao.Update(updateReq)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

}

func (view *AdsView) DeleteAdByID(w http.ResponseWriter, r *http.Request) {

	log.Println(`delete add by id invoked`)

	idStr := r.URL.Query().Get(`id`)

	var id64 int64
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	id64 = int64(id)

	err = view.dao.DeleteAdByID(id64)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)

}
