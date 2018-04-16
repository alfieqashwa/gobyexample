package main

import "fmt"

/*
type Name interface {
	Method1(param_list) return_type
	Method2(param_list) return_type
	...
}
*/
type Information interface {
	General()
	Attributes()
	Inventory()
}

// Implemention of Interface into Concrete Types
type Product struct {
	Name, Description string
	Weight, Price     float64
	Stock             int
}

func (prd Product) General() {
	fmt.Printf("\n%s", prd.Name)
	fmt.Printf("\n%s\n", prd.Description)
	fmt.Println(prd.Price)
}

func (prd Product) Attributes() {
	fmt.Println(prd.Weight)
}

// Type Embedding
type Mobile struct {
	Product
	DisplayFeatures   []string
	ProcessorFeatures []string
}

// Method Overriding
func (mob Mobile) Attributes() {
	mob.Product.Attributes()
	fmt.Println("\nDisplay Features:")
	for _, key := range mob.DisplayFeatures {
		fmt.Println(key)
	}
	fmt.Println("\nProcessor Features:")
	for _, key := range mob.ProcessorFeatures {
		fmt.Println(key)
	}
}
