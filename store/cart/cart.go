package cart

import "github.com/fiiray/packages/store"

type Cart struct {
	CustomName string
	Products   []store.Product
}
func (cart *Cart) GetTotal() (total float64) {
	for _, p := range cart.Products {
		total += p.Price()
	}
	return
}

}