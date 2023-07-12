package ads

import (
	"e2efoo/internal/types"
	"net/http"
)

type Service interface {
	GetAdsByCustomer(customer string) (types.AdResponse, error)
	GetAllAds() (types.AdResponse, error)
}

type API interface {
	GetAds(w http.ResponseWriter, r *http.Request)
}
