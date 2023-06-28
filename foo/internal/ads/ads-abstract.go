package ads

import (
	"e2efoo/internal/types"
	"net/http"
)

type Service interface {
	GetAdsByCustomer(request types.FooRequest) (types.AdResponse, error)
}

type API interface {
	GetAdsByCustomer(w http.ResponseWriter, r *http.Request)
}
