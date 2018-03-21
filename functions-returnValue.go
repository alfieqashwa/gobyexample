package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	rand.Seed(time.Now().Unix())
	var randomValue int

	randomValue = randomWithRange(2, 10)
	fmt.Println("random number:", randomValue)
	randomValue = randomWithRange(2, 10)
	fmt.Println("random number:", randomValue)
	randomValue = randomWithRange(2, 10)
	fmt.Println("random number:", randomValue)
}
func randomWithRange(min, max int) int {
	var value = rand.Int()%(max-min) + min
	return value // block
}

/*

rand.Seed(time.Now().Unix() berada dalam package rand.math dan time.

return value bisa juga dimanfaatkan untuk menghentikan proses dalam blok fungsi dimana ia dipakai.
*/
