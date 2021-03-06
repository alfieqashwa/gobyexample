package main

import "fmt"

type testInt func(int) bool // define a function type of variable

func isOdd(integer int) bool {
	if integer%2 == 0 {
		return false
	}
	return true
}

func isEven(integer int) bool {
	if integer%2 == 0 {
		return true
	}
	return false
}

// pass the function `f` as an argument to another function

func filter(slice []int, f testInt) []int {
	var result []int

	for _, value := range slice {
		if f(value) {
			result = append(result, value)
		}
	}
	return result
}

func main() {
	slice := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	fmt.Println("slice = ", slice)

	odd := filter(slice, isOdd) // use function as values
	fmt.Println("Odd elements of slice are: ", odd)
	even := filter(slice, isEven)
	fmt.Println("Even elements of slice are: ", even)
}
