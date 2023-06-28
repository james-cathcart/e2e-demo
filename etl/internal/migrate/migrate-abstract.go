package migrate

import (
	"e2eetl/internal/types"
	"net/http"
)

type Service interface {
	Migrate(etlRequest types.ETLRequest) (types.ETLResponse, error)
}

type API interface {
	Migrate(w http.ResponseWriter, r *http.Request)
}
