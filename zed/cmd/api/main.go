package main

import (
	"database/sql"
	"e2ezed/internal/ads"
	"e2ezed/internal/ads/view"
	"fmt"
	"log"
	"net/http"

	_ "github.com/mattn/go-sqlite3"
)

const (
	SERVICE_PORT = 8082
)

func main() {

	log.Printf("starting zed service on port -> %d", SERVICE_PORT)

	databaseFile := `ads.db`
	conn, err := sql.Open(`sqlite3`, databaseFile)
	if err != nil {
		log.Fatal(err)
	}

	adDao := ads.New(conn)

	adAPI := view.New(adDao)

	mux := http.NewServeMux()

	mux.HandleFunc(`/ads`, adAPI.GetAds)
	mux.HandleFunc(`/`)

	server := http.Server{
		Addr:    fmt.Sprintf("0.0.0.0:%d", SERVICE_PORT),
		Handler: mux,
	}

	server.ListenAndServe()

}
