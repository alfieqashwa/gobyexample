package main

import "fmt"

func main() {
	for i := 0; i < 5; i++ {
		for j := i; j < 5; j++ {
			fmt.Print(j, " ")
		}

		fmt.Println()
	}
}

/*
=== OUTPUT ===

0 1 2 3 4
1 2 3 4
2 3 4
3 4
4
*/
