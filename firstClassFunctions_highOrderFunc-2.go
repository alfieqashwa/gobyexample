/*
The definition of Higher-order function is a function which does at least one of the following:
    takes one or more functions as arguments
    returns a function as its result

*/
// returning functions from other functions
package main

import (
	"fmt"
)

func simple() func(a, b int) int {
	f := func(a, b int) int {
		return a + b
	}
	return f
}

func main() {
	s := simple()
	fmt.Println(s(60, 7))
}
