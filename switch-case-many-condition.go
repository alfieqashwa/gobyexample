package main

import "fmt"

func main() {

	var point = 6

	switch point {
	case 8:
		fmt.Println("perfect")
	case 7, 6, 5, 4: // sebuah case dapat menampung banyak kondisi
		fmt.Println("awesome")
	default:
		fmt.Println("not bad")
	}
}
