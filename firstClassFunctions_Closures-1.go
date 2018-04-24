/*
Closures are a special case of anonymous functions.
Closures are anonymous functions which access the variables defined outside the body of the function.
*/
package main

import (
	"fmt"
)

func main() {
	a := 5
	func() {
		fmt.Println("A = ", a)
	}()
}
