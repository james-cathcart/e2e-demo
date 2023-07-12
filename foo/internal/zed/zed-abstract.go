package zed

import (
	"e2efoo/internal/types"
)

type Service interface {
	GetAdsByCustomer(customer string) (types.AdResponse, error)
	GetAllAds() (types.AdResponse, error)
}
