package ads

import (
	"e2efoo/internal/types"
	"e2efoo/internal/zed"
	"log"
)

type RestImpl struct {
	zedSvc zed.Service
}

func New(zedSvc zed.Service) Service {
	return &RestImpl{
		zedSvc: zedSvc,
	}
}

func (svc *RestImpl) GetAdsByCustomer(fooRequest types.FooRequest) (types.AdResponse, error) {

	log.Println(`processing ads request`)

	// convert FooRequest to ZedRequest (contrived example)
	zedRequest := types.ZedRequest{}

	// replace empty values with defaults
	if fooRequest.Customer == `` {
		zedRequest.Customer = `all`
	} else {
		zedRequest.Customer = fooRequest.Customer
	}

	adResp, err := svc.zedSvc.GetAdsByCustomer(zedRequest)
	if err != nil {
		return types.AdResponse{}, err
	}

	return adResp, nil
}
