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

func (svc *RestImpl) GetAdsByCustomer(customer string) (types.AdResponse, error) {

	log.Printf("getting ads for customer: %s\n", customer)

	adResp, err := svc.zedSvc.GetAdsByCustomer(customer)
	if err != nil {
		return types.AdResponse{}, err
	}

	return adResp, nil
}

func (svc *RestImpl) GetAllAds() (types.AdResponse, error) {

	log.Println(`getting all ads`)

	adResp, err := svc.zedSvc.GetAllAds()
	if err != nil {
		log.Printf("error: %s\n", err.Error())
		return types.AdResponse{}, err
	}

	return adResp, nil

}
