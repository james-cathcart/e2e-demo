package ads

import (
	"e2ezed/internal/types"
	"net/http"
)

type DAO interface {
	GetAdsByCustomer(customer string) ([]types.Ad, error)
	GetAllAds() ([]types.Ad, error)
}

type API interface {
	GetAds(w http.ResponseWriter, r *http.Request)
}
