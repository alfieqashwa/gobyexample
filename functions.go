package main

import "fmt"

// Return results of a + b and a * b
func SumAndProduct(a, b int) (add int, multiplied int) {
	add, multiplied = a+b, a*b
	return
}

func main() {

	x := 3
	y := 4

	xPLUSy, xTIMESy := SumAndProduct(x, y)

	fmt.Printf("%d + %d = %d\n", x, y, xPLUSy)
	fmt.Printf("%d + %d = %d\n", x, y, xTIMESy)
}
