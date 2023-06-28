package zed

import (
	"bytes"
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

func (svc *RestImpl) GetAdsByCustomer(request types.ZedRequest) (types.AdResponse, error) {

	log.Println(`calling zed API`)

	jsonBytes, err := json.Marshal(request)
	if err != nil {
		return types.AdResponse{}, err
	}

	zedURL := fmt.Sprintf("http://%s/ads", ZedHost)

	httpReq, err := http.NewRequest(
		http.MethodPost,
		zedURL,
		io.NopCloser(bytes.NewReader(jsonBytes)),
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
