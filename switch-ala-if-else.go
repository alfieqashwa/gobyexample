package main

import "fmt"

func main() {

	var point = 6

	switch { //	no variable after switch
	case point == 8: // switch-case ala if-esle
		fmt.Println("perfect")
	case (point < 8) && (point > 3): // switch-case ala-ala if-else
		fmt.Println("awesome")
	default:
		{
			fmt.Println("not bad")
			fmt.Println("you can do better!")
		}
	}
}
