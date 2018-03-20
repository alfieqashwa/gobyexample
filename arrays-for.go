package main

import "fmt"

func main() {
	var fruits = [4]string{"apple", "grape", "banana", "melon"}

	for i := 0; i < len(fruits); i++ {
		fmt.Printf("element %d : %s\n", i, fruits[i])
	}
}

/*
=== OUTPUT ===

element 0 : apple
element 1 : grape
element 2 : banana
element 3 : melon

*/
