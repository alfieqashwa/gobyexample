package main

import (
	"fmt"
)

func SplitValues(f func(sum int) (int, int)) {
	x, y := f(35)
	fmt.Println(x, y) // x = 35 * 5 / 8 == 21, y = 35 - 21 == 14

	x, y = f(50) // x = 50 * 5 / 8 == 31, y = 50 - 31 == 19
	fmt.Println(x, y)
}

func main() {
	a, b := 5, 8
	fn := func(sum int) (int, int) {
		x := sum * a / b
		y := sum - x

		return x, y
	}

	// Passing function value as an argument to another function
	SplitValues(fn)

	// Calling the function value by providing argument
	x, y := fn(20)
	fmt.Println(x, y) // x = 20 * 5 / 8 == 12	| y = 20 - 12 == 8
}
