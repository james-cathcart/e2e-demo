package bar

import (
	"bytes"
	"e2eetl/internal/types"
	"encoding/json"
	"fmt"
	"io"
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

func (svc *RestImpl) Create(records []types.Ad) error {

	barRequest := types.BarRequest{
		Ads: records,
	}

	jsonBytes, err := json.Marshal(barRequest)
	if err != nil {
		return err
	}

	barURL := fmt.Sprintf("http://%s/ads", BarHost)

	request, err := http.NewRequest(
		http.MethodPost,
		barURL,
		io.NopCloser(bytes.NewReader(jsonBytes)),
	)
	if err != nil {
		return err
	}

	response, err := svc.client.Do(request)
	if err != nil {
		return err
	}

	if response.StatusCode != http.StatusCreated {
		return fmt.Errorf("unexpected status code: %d", response.StatusCode)
	}

	return nil
}
