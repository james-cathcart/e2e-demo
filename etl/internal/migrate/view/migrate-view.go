package view

import (
	"e2eetl/internal/migrate"
	"e2eetl/internal/types"
	"encoding/json"
	"io"
	"log"
	"net/http"
)

type MigrateView struct {
	svc migrate.Service
}

func New(svc migrate.Service) migrate.API {
	return &MigrateView{
		svc: svc,
	}
}

func (api *MigrateView) Migrate(w http.ResponseWriter, r *http.Request) {

	log.Println(`migrate api invoked`)

	defer r.Body.Close()

	bodyBytes, err := io.ReadAll(r.Body)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	var etlRequest types.ETLRequest
	err = json.Unmarshal(bodyBytes, &etlRequest)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	etlResponse, err := api.svc.Migrate(etlRequest)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	jsonBytes, err := json.Marshal(etlResponse)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set(`Content-Type`, `application/json`)
	w.WriteHeader(http.StatusCreated)
	w.Write(jsonBytes)
}
