package zed

import (
	"e2efoo/internal/types"
)

type Service interface {
	GetAdsByCustomer(request types.ZedRequest) (types.AdResponse, error)
}
