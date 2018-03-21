package main

import "fmt"

func main() {
	var numbers = []int{2, 3, 0, 4, 3, 5, 2, 0, 4, 0, 3}

	var newNumbers = func(min int) []int {
		var r []int
		for _, e := range numbers {
			if e < min {
				continue
			}
			r = append(r, e)
		}
		return r
	}(3) // IIFE = Immediately-Invoked Function Expression

	fmt.Println("original number:", numbers)
	fmt.Println("filtered number:", newNumbers)

}

/*
=== OUTPUT ===

original number: [2 3 0 4 3 5 2 0 4 0 3]
filtered number: [3 4 3 5 4 3]

*/
