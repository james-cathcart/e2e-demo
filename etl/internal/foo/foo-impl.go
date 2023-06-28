package foo

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

func (svc *RestImpl) GetAdsByCustomer(request types.FooRequest) (types.AdResponse, error) {

	jsonBytes, err := json.Marshal(request)
	if err != nil {
		return types.AdResponse{}, err
	}

	fooURL := fmt.Sprintf("http://%s/ads", FooHost)

	apiRequest, err := http.NewRequest(
		http.MethodPost,
		fooURL,
		io.NopCloser(bytes.NewReader(jsonBytes)),
	)
	if err != nil {
		return types.AdResponse{}, err
	}

	response, err := svc.client.Do(apiRequest)
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
