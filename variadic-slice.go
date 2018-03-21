package main

import "fmt"

func main() {

	var numbers = []int{2, 4, 3, 5, 4, 3, 3, 5, 5, 3} // USING DATA SLICE
	var avg = calculate(numbers...)
	var msg = fmt.Sprintf("Rata-rata : %.2f", avg)

	fmt.Println(msg)
}

func calculate(numbers ...int) float64 {
	var total int = 0
	for _, number := range numbers {
		total += number
	}

	var avg = float64(total) / float64(len(numbers)) // casting: teknik untuk konversi tipe sebuah data ke tipe lain.
	return avg
}

/*
Fungsi Sprintf() pada dasarnya sama dengan fmt.Printf(), hanaya saja fungsi tidak menampilkan nilai, melainkan mengembalikan nilainya dalam bentuk string.
Pada kasus di atas, nilai kembalian fmt.Sprintf() ditampung oleh variable msg .

Selain fmt.Sprintf(), ada juga fmt.Sprint() dan fmt.Sprintln() .

*/
