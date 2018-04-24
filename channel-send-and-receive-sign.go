package main

import (
	"fmt"
	"sync"
)

var wg = sync.WaitGroup{}

func main() {
	ch := make(chan int)
	wg.Add(2)
	go func(ch <-chan int) {
		i := <-ch // receive channel
		fmt.Println(i)
		wg.Done()
	}(ch)
	go func(ch chan<- int) { //send channel
		ch <- 42 // send channel
		wg.Done()
	}(ch)
	wg.Wait()
}
