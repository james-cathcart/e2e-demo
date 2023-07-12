package ads

import (
	"e2ezed/internal/types"
	"net/http"
)

type DAO interface {
	GetAdByID(id int64) ([]types.Ad, error)
	GetAdsByCustomer(customer string) ([]types.Ad, error)
	GetAllAds() ([]types.Ad, error)
	CreateAd(ad types.Ad) (int64, error)
	Update(ad types.Ad) error
	DeleteAdByID(id int64) error
}

type API interface {
	ServeHTTP(w http.ResponseWriter, r *http.Request)
	GetAds(w http.ResponseWriter, r *http.Request)
	CreateAd(w http.ResponseWriter, r *http.Request)
	UpdateAd(w http.ResponseWriter, r *http.Request)
	DeleteAdByID(w http.ResponseWriter, r *http.Request)
}
