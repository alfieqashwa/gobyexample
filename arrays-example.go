package main

import "fmt"

func main() {
	// Declare arrays
	var x [5]int
	// Assign values at specific index
	x[0] = 5
	x[4] = 25
	fmt.Println("Value of x:", x)

	x[1] = 10
	x[2] = 15
	x[3] = 20
	fmt.Println("Value of x:", x)

	// Declare and initialize array with array literal
	y := [5]int{10, 20, 30, 40, 50}
	fmt.Println("Value of y", y)

	// Array literal with ...
	z := [...]int{60, 70, 80, 90, 100}
	fmt.Println("Value of z:", z)
	fmt.Println("Length of z:", len(z))

	// Initialize values at specific index with array literal
	langs := [4]string{0: "Alfie", 3: "Cello"}
	fmt.Println("Value of langs:", langs)
	// Assign values to remaining positions
	langs[1] = "Bunda"
	langs[2] = "Qashwa"

	// Iterate over the elements of array
	fmt.Println("Values of langs:", langs)
	fmt.Println("\nIterate over arrays\n")
	for i := 0; i < len(langs); i++ {
		fmt.Printf("langs[%d]:%s \n", i, langs[i])
	}
	fmt.Println("\n")

	//Interate over the elements of array using range
	for k, v := range langs {
		fmt.Printf("langs[%d]:%s \n", k, v)
	}
}
