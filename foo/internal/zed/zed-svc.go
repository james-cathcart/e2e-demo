package zed

import (
	"e2efoo/internal/types"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
)

type RestImpl struct {
	client *http.Client
}

func New() Service {
	return &RestImpl{
		client: &http.Client{},
	}
}

func (svc *RestImpl) GetAdsByCustomer(customer string) (types.AdResponse, error) {

	log.Println(`calling zed API`)

	zedURL := fmt.Sprintf("http://%s/ads?filter=customer&value=%s", ZedHost, customer)

	httpReq, err := http.NewRequest(
		http.MethodGet,
		zedURL,
		nil,
	)
	if err != nil {
		return types.AdResponse{}, err
	}

	response, err := svc.client.Do(httpReq)
	if err != nil {
		return types.AdResponse{}, err
	}

	defer response.Body.Close()
	bodyBytes, err := io.ReadAll(response.Body)
	if err != nil {
		return types.AdResponse{}, err
	}

	var adResponse types.AdResponse
	err = json.Unmarshal(bodyBytes, &adResponse)
	if err != nil {
		return types.AdResponse{}, err
	}

	return adResponse, nil

}

func (svc *RestImpl) GetAllAds() (types.AdResponse, error) {

	log.Println(`GetAllAds invoked`)

	zedURL := fmt.Sprintf("http://%s/ads", ZedHost)

	httpReq, err := http.NewRequest(
		http.MethodGet,
		zedURL,
		nil,
	)
	if err != nil {
		log.Printf("error: %s\n", err.Error())
		return types.AdResponse{}, err
	}

	response, err := svc.client.Do(httpReq)
	if err != nil {
		log.Printf("error: %s\n", err.Error())
		return types.AdResponse{}, err
	}

	defer response.Body.Close()
	bodyBytes, err := io.ReadAll(response.Body)
	if err != nil {
		log.Printf("error: %s\n", err.Error())
		return types.AdResponse{}, err
	}

	var adResponse types.AdResponse
	err = json.Unmarshal(bodyBytes, &adResponse)
	if err != nil {
		log.Printf("error: %s\n", err.Error())
		return types.AdResponse{}, err
	}

	return adResponse, nil

}
