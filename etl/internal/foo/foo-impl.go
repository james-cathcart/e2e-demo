package foo

import (
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

func (svc *RestImpl) GetAdsByCustomer(customer string) (types.AdResponse, error) {

	fooURL := fmt.Sprintf("http://%s/ads?filter=customer&value=%s", FooHost, customer)

	apiRequest, err := http.NewRequest(
		http.MethodGet,
		fooURL,
		nil,
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
