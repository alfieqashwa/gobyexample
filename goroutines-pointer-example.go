package main

import (
	"fmt"
	"sync"
)

func main() {
	var wg sync.WaitGroup
	var msg = "Hello"

	wg.Add(1)
	go sayHello(msg, &wg)
	wg.Wait()
}

func sayHello(msg string, wg *sync.WaitGroup) {
	fmt.Println(msg)
	wg.Done()
}
