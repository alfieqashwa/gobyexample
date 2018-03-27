package main

import "fmt"

type Human struct {
	name  string
	age   int
	phone string
}

type Student struct {
	Human  // anonymous field
	school string
}

type Employee struct {
	Human
	company string
}

func (h *Human) SayHi() {
	fmt.Printf("Hi, I am %s you can call me on %s\n", h.name, h.phone)
}

func (e *Employee) SayHi() {
	fmt.Printf("Hi, my name is %s, I work at %s. Call me on %s\n", e.name, e.company, e.phone)
}

func main() {

	cello := Student{Human{"Cello", 25, "222-222-YYYY"}, "MIT"}
	qashwa := Employee{Human{"Qashwa", 35, "111-888-XXXX"}, "Golang Inc"}

	cello.SayHi()
	qashwa.SayHi()
}
