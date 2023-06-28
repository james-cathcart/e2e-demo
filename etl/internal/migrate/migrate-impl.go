package migrate

import (
	"e2eetl/internal/bar"
	"e2eetl/internal/foo"
	"e2eetl/internal/types"
)

type SimpleImpl struct {
	fooSvc foo.Service
	barSvc bar.Service
}

func New(fooSvc foo.Service, barSvc bar.Service) Service {
	return &SimpleImpl{
		fooSvc: fooSvc,
		barSvc: barSvc,
	}
}

func (svc *SimpleImpl) Migrate(etlRequest types.ETLRequest) (types.ETLResponse, error) {

	fooRequest := types.FooRequest{
		Customer: etlRequest.Customer,
	}

	fooRecords, err := svc.fooSvc.GetAdsByCustomer(fooRequest)
	if err != nil {
		return types.ETLResponse{}, err
	}

	if !etlRequest.DryRun {
		err = svc.barSvc.Create(fooRecords.Ads)
		if err != nil {
			return types.ETLResponse{}, err
		}
	}

	// return list of migrated records
	etlResponse := types.ETLResponse{
		Ads:      fooRecords.Ads,
		Migrated: len(fooRecords.Ads),
	}

	return etlResponse, nil

}
