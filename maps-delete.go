package main

import "fmt"

func main() {
	var chicken = map[string]int{"january": 50, "february": 40}

	fmt.Println(len(chicken)) // 2
	fmt.Println(chicken)

	delete(chicken, "january")

	fmt.Println(len(chicken)) // 1
	fmt.Println(chicken)

}
