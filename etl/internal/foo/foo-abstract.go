package foo

import "e2eetl/internal/types"

type Service interface {
	GetAdsByCustomer(request types.FooRequest) (types.AdResponse, error)
}
