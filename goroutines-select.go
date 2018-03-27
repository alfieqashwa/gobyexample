/*
How can we deal with more than one channel?
Go has a keyword called select to listen to many channels

Select is blocking by default and it continues to execute only when one of channels has data to send or receive.
If several channels are ready to use at the same time, select chooses which to execute randomly.
*/
package main

import "fmt"

func fibonacci(c, quit chan int) {
	x, y := 1, 1

	for {

		select {
		case c <- x:
			x, y = y, x+y
		case <-quit:
			fmt.Println("quit")
			return
		}
	}
}

func main() {
	c := make(chan int)
	quit := make(chan int)
	go func() {
		for i := 0; i < 10; i++ {
			fmt.Println(<-c)
		}
		quit <- 0
	}()
	fibonacci(c, quit)
}

/*
=== OUTPUT ===

1
1
2
3
5
8
13
21
34
55
quit

*/
