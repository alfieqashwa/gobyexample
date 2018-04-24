/*
The definition of Higher-order function is a function which does at least one of the following:
    takes one or more functions as arguments
    returns a function as its result

*/
// passing functions as arguments to other functions
package main

import (
	"fmt"
)

func simple(a func(a, b int) int) {
	fmt.Println(a(60, 7))
}

func main() {
	f := func(a, b int) int {
		return a + b
	}
	simple(f)
}
