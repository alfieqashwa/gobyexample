package main

import "fmt"

func main() {
	var fruits = [4]string{"apple", "grape", "banana", "melon"}

	// for i := 0; i < len(fruits); i++ {
	for i, fruit := range fruits { // for - range
		fmt.Printf("element %d : %s\n", i, fruit)
	}
}

/*
=== OUTPUT ===

element 0 : apple
element 1 : grape
element 2 : banana
element 3 : melon

Array fruits diambil elemen-nya secara berurutan.
Nilai tiap elemen ditampung variable oleh fruit (tanpa huruf s).
Sedangkan indeks nya ditampung variable i.
*/
