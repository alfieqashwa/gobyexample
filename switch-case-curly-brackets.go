package main

import "fmt"

func main() {

	var point = 3

	switch point {
	case 8:
		fmt.Println("perfect")
	case 7, 6, 5, 4: // sebuah case dapat menampung banyak kondisi
		fmt.Println("awesome")
	default:
		{
			fmt.Println("not bad")
			fmt.Println("you can do better!")
		}
	}
}

/*
Tand kurung kurawal bisa titerapkan pada keyword case dan default. Opsional, boleh dipakai boleh tidak. Baik dipakai jika pada blok kondisi yang di dalamnya ada banyak statement agar kode terlihat lebih rapi dan mudah di maintain.
*/
