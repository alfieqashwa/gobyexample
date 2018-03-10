/*
Variadic Function == pembuatan fungsi dengan parameter sejenis yang tak terbatas.
Maksudnya tak terbatas disini adalah jumlah parameter yang disisipkan ketika pemanggilan fungsi bisa berapa saja.

Parameter variadic memiliki sifat yang mirip dengan slice. Nilai dari parameter-parameter yang disisipkan bertipe data sama,
dan ditampung oleh sebuah variable saja.
Cara pengaksesan tiap datanya juga sama, dengan menggunakan index.

Deklarasi parameter variadic sama dengan cara deklarasi variabel biasa, pembedanya adalah pada parameter jenis ini ditambahkan tanda titik tiga kali ( ... ) tepat setelah penulisan variable (sebelum tipe data).
Nantinya semua nilai yang disisipkan sebagai parameter akan ditampung oleh variable tersebut.
*/

/*Variadic functions can be called with any number of trailing arguments.
For example, fmt.Println is a common variadic function. */
package main

import "fmt"

// Here’s a function that will take an arbitrary number of ints as arguments.
func sum(nums ...int) {
	fmt.Print(nums, " ")
	total := 0
	for _, num := range nums {
		total += num
	}
	fmt.Println(total)
}

func main() {

	// Variadic functions can be called in the usual way with individual arguments.
	sum(1, 2)
	sum(1, 2, 3)

	// If you already have multiple args in a slice, apply them to a variadic function using func(slice...) like this.
	nums := []int{1, 2, 3, 4}
	sum(nums...)
}

/*
$ go run variadic-functions.go

[1 2] 3
[1 2 3] 6
[1 2 3 4] 10

*/
