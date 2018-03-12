package main

import "fmt"

func Add(x, y int) float32 {
	return float32(x + y)
}

func Subtract(x, y int) float32 {
	return float32(x - y)
}

func main() {
	x, y := 50, 15
	hasil := Add(x, y)
	fmt.Printf("Hasil dari %d ditambah %d adalah %.1f\n", x, y, hasil) // Hasil dari 50 ditambah 15 adalah 65.0

	hasil = Subtract(x, y)
	fmt.Printf("Hasil dari %d dikurang %d adalah %.1f\n", x, y, hasil) // Hasil dari 50 di kurang 15 adalah 35.0
}
