/*
Let's conclude this section by writing one more program.
This program will perform the same operations on each element of a slice and return the result.
For example if we want to multiply all integers in a slice by 5 and return the output,
it can be easily done using first class functions.
These kind of functions which operate on every element of a collection are called map functions.
I have provided the program below. It is self explanatory.
*/
package main

import "fmt"

func iMap(s []int, f func(int) int) []int {
	var r []int
	for _, v := range s {
		r = append(r, f(v))
	}
	return r
}
func main() {
	a := []int{5, 6, 7, 8, 9}
	r := iMap(a, func(n int) int {
		return n * 5
	})

	sum := 0
	for _, v := range r {
		sum += v
	}
	fmt.Printf("Sum of %d is %d\n", r, sum)
}
