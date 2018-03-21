package main

import (
	"fmt"
	"math"
)

func main() {
	var diameter float64 = 15
	var area, circumference = calculate(diameter)

	fmt.Printf("Luas Lingkaran\t\t: %.2f \n", area)
	fmt.Printf("Keliling Lingkaran\t: %.2f \n", circumference)
}

func calculate(d float64) (area, circumference float64) { // Predefined Return Value
	area = math.Pi * math.Pow(d/2, 2)
	circumference = math.Pi * d

	return // just return
}
