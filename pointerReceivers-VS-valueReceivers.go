package main

import (
	"fmt"
)

type Employee struct {
	name string
	age  int
}

func (e Employee) changeName(newName string) {
	e.name = newName
}

func (e *Employee) changeAge(newAge int) {
	e.age = newAge
}

func main() {
	e := Employee{
		name: "Mark Andrew",
		age:  50,
	}
	fmt.Printf("Employee name before change: %s", e.name)
	e.changeName("Michael Knight") // never happen
	fmt.Printf("\nEmployee name after change %s", e.name)

	fmt.Printf("\n\nEmployee age before change %d", e.age)
	(&e).changeAge(67) // or just write e.changeAge(67)
	fmt.Printf("\nEmployee age after change %d", e.age)
}
