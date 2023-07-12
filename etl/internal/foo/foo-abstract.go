package foo

import "e2eetl/internal/types"

type Service interface {
	GetAdsByCustomer(customer string) (types.AdResponse, error)
}
