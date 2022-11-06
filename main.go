package main

import (
	"fmt"
	. "github.com/fiiray/packages/fmt"
	"github.com/fiiray/packages/store"
	"github.com/fiiray/packages/store/cart"
)

func main() {
	product := store.NewProduct("Kayak", "Watersports", 279)

	cart := cart.Cart{
		CustomName: "Alice",
		Products:   []store.Product{*product},
	}

	fmt.Println("Name:", cart.CustomName)
	fmt.Println("Total:", ToCurrency(cart.GetTotal()))
}
