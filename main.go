package main

import (
	"fmt"
	"github.com/fiiray/packages/store"
)

func main() {
	product := store.Product{
		Name:     "Shoes",
		Category: "Clothing",
	}

	fmt.Println("Name:", product.Name)
	fmt.Println("Category:", product.Category)

}