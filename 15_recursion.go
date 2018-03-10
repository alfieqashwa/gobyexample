// Go supports recursive functions. Here’s a classic factorial example.
package main

import "fmt"

// This fact function calls itself until it reaches the base case of fact(0).
func fact(n int) int {
	if n == 0 {
		return 1
	}
	return n * fact(n-1)
}

func main() {
	fmt.Println(fact(1))
	fmt.Println(fact(2))
	fmt.Println(fact(3))
	fmt.Println(fact(4))
	fmt.Println(fact(5))
	fmt.Println(fact(6))
	fmt.Println(fact(7))
}

// i'm still don't understand!
