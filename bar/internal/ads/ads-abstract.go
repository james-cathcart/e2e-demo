package ads

import (
	"e2ebar/internal/types"
	"net/http"
)

type Service interface {
	Create(records ...types.Ad)
}

type API interface {
	Create(w http.ResponseWriter, r *http.Request)
}
