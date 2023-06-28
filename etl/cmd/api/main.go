package main

import (
	"e2eetl/internal/bar"
	"e2eetl/internal/foo"
	"e2eetl/internal/migrate"
	"e2eetl/internal/migrate/view"
	"fmt"
	"log"
	"net/http"
)

const (
	SERVICE_PORT = 8084
)

func main() {

	log.Printf("starting etl service on port -> %d", SERVICE_PORT)

	fooSvc := foo.New()
	barSvc := bar.New()
	migrateSvc := migrate.New(fooSvc, barSvc)
	migrateAPI := view.New(migrateSvc)

	mux := http.NewServeMux()
	mux.HandleFunc(`/etl`, migrateAPI.Migrate)

	server := http.Server{
		Addr:    fmt.Sprintf("0.0.0.0:%d", SERVICE_PORT),
		Handler: mux,
	}

	server.ListenAndServe()

}
