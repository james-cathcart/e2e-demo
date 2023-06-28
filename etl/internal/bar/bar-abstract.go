package bar

import "e2eetl/internal/types"

type Service interface {
	Create(records []types.Ad) error 
}
