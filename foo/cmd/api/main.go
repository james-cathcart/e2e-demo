package main

import (
	"e2efoo/internal/ads"
	"e2efoo/internal/ads/view"
	"e2efoo/internal/zed"
	"fmt"
	"log"
	"net/http"
)

const (
	SERVICE_PORT = 8081
)

func main() {

	log.Printf("starting foo service on port -> %d", SERVICE_PORT)

	zedSvc := zed.New()
	adsSvc := ads.New(zedSvc)
	adsAPI := view.NewAdsAPI(adsSvc)

	mux := http.NewServeMux()

	mux.HandleFunc(`/ads`, adsAPI.GetAds)

	server := http.Server{
		Addr:    fmt.Sprintf("0.0.0.0:%d", SERVICE_PORT),
		Handler: mux,
	}

	server.ListenAndServe()

	log.Println(`foo service exiting...`)

}
