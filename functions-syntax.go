/*
=== FUNCTION SYNTAX ===

func funcName(input1 type1, input2 type2) (output1 type1, output2 type2) {

	// function body
	// multi-value return
	return value1, value2

}

=== EXPLAINATION ===
Functions may have zero, one or more than one arguments. The argument type comes after
the argument name and arguments are separated by , .
Functions can return multiple values.
There are two return values named output1 and output2 , you can omit their names and use
their type only.
If there is only one return value and you omitted the name, you don’t need brackets for the
return values.
If the function doesn’t have return values, you can omit the return parameters altogether.
If the function has return values, you have to use the return statement somewhere in the
body of the function.
*/
package main

import f "fmt"

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func main() {
	x, y, z := 3, 4, 5

	max_xy := max(x, y) //	call function max(x, y)
	max_xz := max(x, z) //	call function max(x, z)

	f.Printf("max(%d, %d) = %d\n", x, y, max_xy)
	f.Printf("max(%d, %d) = %d\n", x, z, max_xz)
	f.Printf("max(%d, %d) = %d\n", y, z, max(y, z)) // call function here

}
