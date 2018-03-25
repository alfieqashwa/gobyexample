/*
=== OBJECT-ORIENTED ===

Object oriented languages allow programmers to declare a function inside the class definition. Go doesn't allow us to do that, we have to declare a method on a struct via a special syntax.

this is a NOT-RECOMMENDED example code in golang programming language

*/
package main

import "fmt"

type Rectangle struct {
	width, height float64
}

func area(r Rectangle) float64 {
	return r.width * r.height
}

func main() {
	r1 := Rectangle{12, 2}
	r2 := Rectangle{9, 4}

	fmt.Println("Area of r1 is ", area(r1))
	fmt.Println("Area of r2 is ", area(r2))
}
