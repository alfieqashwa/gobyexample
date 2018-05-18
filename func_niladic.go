package main

import (
	"fmt"
	"time"
)

// Declares a func named nop
func nop() {
	// here it does nothing
}

func tick() {
	fmt.Println(time.Now().Format(time.Kitchen))
}

func main() {
	// it does nothing and doesn't return anything
	nop()
	nop()
	nop()
	nop()

	// prints the time and doesn't return anything
	tick()
}
