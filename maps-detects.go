package main

import "fmt"

func main() {
	var chicken = map[string]int{"january": 50, "february": 40}
	var value, isExist = chicken["may"]

	if isExist {
		fmt.Println(value)
	} else {
		fmt.Println("item is not exists")
	}
}
