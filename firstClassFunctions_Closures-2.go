/*
Closures are a special case of anonymous functions.
Closures are anonymous functions which access the variables defined outside the body of the function.
*/

// Every closure is bound to its own surrounding variable. Let's understand what this means by using a simple example.
package main

import (
	"fmt"
)

func appendStr() func(string) string {
	t := "Hello"
	c := func(b string) string {
		t += " " + b // t = t + " " + b
		return t
	}
	return c
}

func main() {
	a := appendStr()
	b := appendStr()
	fmt.Println(a("World"))
	fmt.Println(b("Ev'ryone"))

	fmt.Println(a("Gopher"))
	fmt.Println(b("!"))
}
