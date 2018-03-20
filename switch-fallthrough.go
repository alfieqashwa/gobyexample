package main

import "fmt"

func main() {

	var point int = 6

	switch {
	case point == 8:
		fmt.Println("perfect")
	case (point < 8) && (point > 3):
		fmt.Println("awesome")
		fallthrough // memaksa program untuk mengecek lebih lanjut
	case point < 5:
		fmt.Println("you need to learn more")
	default:
		{
			fmt.Println("not bad")
			fmt.Println("be better, dude!")
		}
	}
}

/*
Switch pada Golang memiliki perbedaan dengan bahasa lain; manakala sebuah case terpenuhi, pengecekan kondisi tidak akan diteruskan ke case-case setelahnya.
Keyword fallthrough digunakan untuk memaksa proses pengecekkan diteruskan ke case selanjutnya.

=== OUTPUT ===
awesome
you need to learn more
*/
