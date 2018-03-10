/*
Pointer adalah reference atau alamat memory. Varibale pointer berarti variable yang berisi alamat memori suatu nilai.
Sebagai contoh sebuah variable bertipe integere memiliki nilai 4, maka yang dimaksud pointer adalah "alamat memori dimiana nilai 4 di simpan", bukan nilai 4 nya sendiri.

Variable-variable yang memiliki reference atau alamat pointer yang sama, saling berhubungan satu sama lain dan nilainya pasti sama.
Ketika ada perubahan nilai, maka akan memberikan efek kepada variable lain( yang referensi-nya sama) yaitu nilainya ikut berubah.

*/
package main

import "fmt"

func main() {

	var numberA int = 4
	var numberB *int = &numberA

	fmt.Println("numberA (value)	:", numberA)
	fmt.Println("numberA (address)	:", &numberA)

	fmt.Println("numberB (value)	:", *numberB)
	fmt.Println("numberB (address)	:", numberB)
}
