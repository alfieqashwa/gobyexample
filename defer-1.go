package main

import "fmt"

func main() {
	defer fmt.Println("The third line.")
	defer fmt.Println("The second line.")
	defer fmt.Println("The first line.")
}
