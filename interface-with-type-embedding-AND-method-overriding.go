// Example of Interface with Type Embedding and Method Overriding on GO language
package main

import (
	"fmt"
)

/*
The interface type Information is a contract tof creating
various Product types in a catalog. The information interface
 provide three behaviours in its contract: General, Attributes, and Inventory.
*/
type Information interface {
	General()
	Attributes()
	Inventory()
}

/*
A struct Product is declared with fields for holding its state and methods
implemented based on the behaviours defined in the Information interface. */
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

// A struct Mobile is declared in which the type Product is embedded
type Mobile struct {
	Product
	DisplayFeatures   []string
	ProcessorFeatures []string
}

// Override for the Attributes Method for the Mobile struct
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

// A struct Shirts is declared in which the type Product is embedded.
type Shirts struct {
	Product
	Size, Pattern, Sleeve, Fabric string
}

// Overrides for the Attributes Method for the Shirts struct
func (shrt Shirts) Attributes() {
	fmt.Println("\nSpecification:")
	fmt.Println(shrt.Size)
	fmt.Println(shrt.Pattern)
	fmt.Println(shrt.Sleeve)
	fmt.Println(shrt.Fabric)
}

func (prd Product) Inventory() {
	fmt.Println(prd.Stock)
}

// A struct named Catalog is declared to represent catalog of various products.
// The type of Details foeld is used a slice of the Information interface
type Catalog struct {
	Name    string
	Details []Information
}

func (c Catalog) CatalogDetails() {
	fmt.Println("**************************************************")
	fmt.Printf("\nStore : %s \n", c.Name)
	for _, key := range c.Details {
		fmt.Println("===================Products==================")
		key.General()
		key.Attributes()
		key.Inventory()
		fmt.Println("=============================================")
	}
}

func main() {
	//Create an Instance of Mobile Type and Call Methods
	MobilePrd := Mobile{
		Product{
			"Apple iPhone 6s (Rose Gold, 32 GB)		(2 GB RAM)",
			"Handset, Apple EarPods with Remote and Mic, Lightning to USB Cable",
			40999, 143, 10,
		},
		[]string{"2 GB RAM", "4.7 inch Retina HD Display", "12MP Primary Camera"},
		[]string{"32 GB RAM", "4.7 inch Retina HD Display", "5MP Primary Font"},
	}

	//Create an Instance of Shirts Type and Call Meethods
	shirtPrd := Shirts{
		Product{
			"Reebok Stripped Men's Round Neck Grey T-Shirt",
			"Machine Wash as per Tag, Do Not Dry Clean, Do Not Bleach",
			1124, 400, 25,
		},
		"XL", "Stripped", "Half Sleeve", "Cotton",
	}

	//Create an Instance of Catalog Type
	category := Catalog{
		"Amazon",
		[]Information{MobilePrd, shirtPrd},
	}
	category.CatalogDetails()

}
