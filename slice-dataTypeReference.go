/*
Slice merupakan tipe data reference. Artinya jika ada slice baru yang terbentuk dari slice lama, maka data elemen slice yang baru akan memiliki alamat memori yang sama dengan elemen slice lama.
Setiap perubahan yang terjadi di elemen slice baru, akan berdampak juga pada elemen slice lama yang memiliki referensi yang sama.

Let's prove it!!
*/
package main

import "fmt"

func main() {
	var bands = []string{
		"radiohead",
		"nirvana",
		"muse",
		"coldplay",
	}

	var aBands = bands[0:3]
	var bBands = bands[1:4]

	var aaBands = aBands[1:2]
	var bbBands = bBands[0:1]

	fmt.Println(bands)   // [radiohead nirvana muse coldplay]
	fmt.Println(aBands)  // [radiohead nirvana muse]
	fmt.Println(bBands)  // [nirvana muse coldplay]
	fmt.Println(aaBands) // [nirvana]
	fmt.Println(bbBands) // [nirvana]

	// Band "nirvana" di ubah menjadi "queen"
	bbBands[0] = "queen"

	fmt.Println(bands)   // [radiohead queen muse coldplay]
	fmt.Println(aBands)  // [radiohead queen muse]
	fmt.Println(bBands)  // [queen muse coldplay]
	fmt.Println(aaBands) // [queen]
	fmt.Println(bbBands) // [queen]

}
