package main

import (
	"fmt"
	. "github.com/fiiray/packages/fmt"
	"github.com/fiiray/packages/store"
)

func main() {
	product := store.NewProduct("Kayak", "Watersports", 279)
	fmt.Println("Name:", product.Name)
	fmt.Println("Category:", product.Category)
	fmt.Println("Price:", ToCurrency(product.Price()))
}
