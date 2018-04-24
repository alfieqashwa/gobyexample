package main

import "fmt"

func main() {
	a := func() {
		fmt.Println("Hello, First Class Function!")
	}
	a()
	fmt.Printf("%T", a)
}
