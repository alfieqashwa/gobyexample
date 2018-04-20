package main

import "fmt"

type rectangle struct {
	length int
	width  int
}

func perimeter(r *rectangle) { // pointer arguments in functions
	fmt.Println("Perimeter Function output:", 2*(r.length+r.width))
}

func (r *rectangle) perimeter() { // pointer receivers in methods
	fmt.Println("Perimeter Method output:", 2*(r.length+r.width))
}

func main() {
	r := rectangle{
		length: 10,
		width:  5,
	}

	p := &r //pointer to r
	perimeter(p)
	p.perimeter()

	/*
		cannot use r (type rectangle) as type *rectangle in arguments to perimeter
		in argument to area
	*/
	// perimeter(r)

	p.perimeter() // calling pointer receiver with a value
}
