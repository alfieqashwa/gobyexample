package main

import "fmt"

func main() {
	func() {
		fmt.Println("Hello, First Class Function!")
	}()
}
