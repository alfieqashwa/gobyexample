package main

func main() {
	a, b := 1, 0 // panic will happens as dividing an integer by zero
	_ = a / b
}
