package main

import (
	"fmt"
)

type Human struct {
	name  string
	age   int
	phone string // Human has phone field
}

type Employee struct {
	Human     // embedded field Human
	specialty string
	phone     string // phone in employee
}

func main() {

	Bob := Employee{Human{"Bob", 34, "0217490001"}, "Designer", "081234567879"}

	fmt.Println("Bob's work phone is:", Bob.phone)

	// access phone field in Human
	fmt.Println("Bob's personal phone is:", Bob.Human.phone)
}
