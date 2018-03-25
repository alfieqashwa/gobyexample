/*
=== OBJECT-ORIENTED in GOLANG ===

"A method is a function with an implicit first argument, called a receiver" -Rob Pike

=== SYNTAX of METHOD ===

func (r ReceiverType) funcName(parameters) (results)

*/

package main

import (
	"fmt"
	"math"
)

type Rectangle struct {
	width, height float64
}
type Circle struct {
	radius float64
}

func (r Rectangle) area() float64 {
	return r.width * r.height
}

func (c Circle) area() float64 {
	return c.radius * c.radius * math.Pi
}

func main() {

	r1 := Rectangle{12, 2}
	r2 := Rectangle{9, 4}
	c1 := Circle{10}
	c2 := Circle{25}

	fmt.Println("Area of r1 is: ", r1.area())
	fmt.Println("Area of r2 is: ", r2.area())
	fmt.Printf("Area of c1 is: %.2f\n", c1.area())
	fmt.Printf("Area of c2 is: %.2f\n", c2.area())
}

/*
=== NOTES ===

If the bane of methods are same but they don't share the same receivers, they are not the same.
Methods are able to access fields within receivers.
Use . (dot) to call a method in the struct, the same way fields are called.

One thing that's worth noting is that the method with a dotted line means the receiver is passed by value, not by reference.
The difference between them is that a method can change its receiver's values when the receiver is passed by reference, and it gets a copy of the receiver when the receiver is passed by value.

Any type can be the receiver of the method.
*/
