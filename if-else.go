package main

import "fmt"

func main() {

	var nilai int = 8

	if nilai == 10 {
		fmt.Println("Lulus sempurna")
	} else if nilai > 5 {
		fmt.Println("Lulus")
	} else {
		fmt.Println("Tidak lulus")
	}
}
