package main

import "fmt"

func main() {

	var kondisi = "kuning"
	var point = 6

	switch kondisi {
	case "merah":
		fmt.Println("Berhenti!")
	case "kuning":
		fmt.Println("Siap-siap")
	default:
		fmt.Println("Jalan!")
	}

	switch point {
	case 8:
		fmt.Println("perfect")
	case 7:
		fmt.Println("awesoome")
	default:
		fmt.Println("not bad")
	}
}

/*
Dalam bahasa pemgoraman Golang, manakala sebuah case terpenuhi, program  tidak akan melanjutkan pengecekan case selanjutnya meskipun tidak ada keyword "break", berbeda dengan bahasa pemograman lain.
*/
