// Functions are central in Go. We’ll learn about functions with a few different examples.

package main

import "fmt"

// Here’s a function that takes two ints and returns their sum as an int.
func plus(a int, b int) int {

	// Go requires explicit returns, i.e. it won’t automatically return the value of the last expression
	return a + b
}

/* When you have multiple consecutive parameters of the same type,
you may omit the type name for the like-typed parameters up to the final parameter that declares the type. */
func plusPlus(a, b, c int) int {
	return a + b + c
}

// Call a function just as you’d expect, with name(args).
func main() {

	res := plus(1, 2)
	fmt.Println("5+6 =", res)

	res = plusPlus(4, 5, 9)
	fmt.Println("4+5+9 =", res)
	// see the "=", not ":=" Why??, i don't know yet.
	// Now I know why. Because i usedd the same variable, which is var res.
	// If I create new var, let say, result := plusPlus(4,5,9) that works!
	//The error msg: """./functions.go:26: no new variables on left side of :="""
}

/*
go run functions.go

5+6 = 3
4+5+9 = 18

*/
