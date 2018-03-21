package main

import "fmt"
import "strings"

func main() {

	var names1 = []string{"Cello", "el", "Zhafran"}
	printMessage("Hola,", names1)

	var names2 = []string{"Qashwa", "Zhafira"}
	printMessage("Hello,", names2)

}

func printMessage(message string, arr []string) { // two parameters: message and slice arr

	var nameString = strings.Join(arr, " ")
	fmt.Println(message, nameString)

}

/*
2 buah parameter dipanggil pada main function:
	1. "Hola," yang ditampung parameter message, dan
	2. variable names1 & names2 yang ditampung oleh parameter arr .

Dalam printMessage function, nilai arr yg merupakan slice string digabungkan menjadi sebuah string dengan pembatasnya adalah karakter spasi (" ").

Penggabungan slice dapat dilakukan dengan memanfaatkan fungsi strings,Join() yang berada di dalam package strings (see import "strings")

*/
