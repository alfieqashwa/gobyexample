package main

import (
	"fmt"
)

func main() {
qashwaLoop: // Nama label dapat diganti dengan nama lain (contoh: outerLoop) dan harus diakhiri dengan tanda titik dua ( : )
	for i := 0; i < 5; i++ {
		for j := 0; j < 5; j++ {
			if i == 3 {
				break qashwaLoop
			}
			fmt.Print("matriks [", i, "][", j, "]", "\n")
		}
	}
}

/*
=== OUTPUT ===
matriks [0][0]
matriks [0][1]
matriks [0][2]
matriks [0][3]
matriks [0][4]
matriks [1][0]
matriks [1][1]
matriks [1][2]
matriks [1][3]
matriks [1][4]
matriks [2][0]
matriks [2][1]
matriks [2][2]
matriks [2][3]
matriks [2][4]
*/
