package main

import (
	"e2ebar/internal/ads"
	"e2ebar/internal/ads/view"
	"fmt"
	"log"
	"net/http"
)

const (
	SERVICE_PORT = 8083
)

func main() {

	log.Printf("starting bar service on port -> %d", SERVICE_PORT)

	adsSvc := ads.New()
	adsView := view.New(adsSvc)

	mux := http.NewServeMux()

	mux.HandleFunc(`/ads`, adsView.Create)

	server := http.Server{
		Addr:    fmt.Sprintf("0.0.0.0:%d", SERVICE_PORT),
		Handler: mux,
	}

	server.ListenAndServe()

}
