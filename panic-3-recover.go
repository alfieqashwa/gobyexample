package main

import (
	"fmt"
)

func main() {
	fmt.Println("hi!")
	defer func() {
		v := recover()
		fmt.Println("recovered:", v)
	}()
	panic("bye!")
	fmt.Println("unreachable")
}

/*
	the output:
	hi!
	recovered: bye!
*/
