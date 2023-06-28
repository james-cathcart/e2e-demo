package ads

import (
	"e2ebar/internal/types"
	"fmt"
)

type ConsoleImpl struct {
}

func New() Service {
	return &ConsoleImpl{}
}

func (svc *ConsoleImpl) Create(ads ...types.Ad) {

	fmt.Println("\nadding records:\n")
	fmt.Println("Customer       Info")
	fmt.Println("---------      -----")

	for _, ad := range ads {
		fmt.Printf("%-14s %s\n", ad.Customer, ad.Info)
	}

	fmt.Println()

}
