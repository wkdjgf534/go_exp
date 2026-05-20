package main

import (
	"fmt"
	"time"

	"02-factory/products"
)

func main() {

	// With Factory pattern
	factory := products.Product{}
	product1 := factory.New()
	fmt.Println("My first product was created at", product1.CreatedAt.UTC())
	fmt.Println(product1)

	// ------------------------------------

	// Without Factory pattern
	product2 := products.Product{
		ProductName: "widget",
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
	}
	fmt.Println("My second product was created at", product1.CreatedAt.UTC())
	fmt.Println(product2)
}
