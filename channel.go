package main

import (
	"fmt"
	"time"
)

func printCount(c chan int) { // (1) An int type channel passed in
	num := 0
	for num >= 0 {
		num = <-c // (2) Waits for value to come in
		fmt.Print(num, " ")
	}
}

func main() {
	c := make(chan int) // (3) A channel is created
	a := []int{8, 6, 7, 5, 3, 0, 9, -1}

	go printCount(c) // (4) Starts the goroutine

	for _, v := range a {
		c <- v // (5) Passes ints into channel
	}
	time.Sleep(time.Millisecond * 1) // (6) main pauses before ending.
	fmt.Println("End of main")
}

/*
=== EXPLANATION ===

At the start of main, an integer-typed channel c is created (3) to communicate between goroutines. When printCount is started as a goroutine, the channel is passed in (4). As an argument to printCount, the channel needs to be identified as an integer channel (1). In the for loop inside printCount, num waits for channel c to send in integers (2).
Back in main, a list of integers is iterated over and passed into the channel c one at a time (5). When each integer is passed into the channel on main (5), it's received into num within printCount (2). printCount continues until the for loop goes into another iteration and comes to the channel statement again (2), where it waits for another value to come in on the channel. After main is done iterating over the integers, it continues on.
When main is finished executing, the entire program is done, so you pause for a second (6) before exiting so printCount can complete before main is done.
Running this code produces the following listing.

=== Channel Ouput ===

8 6 7 5 3 0 9 -1 End of main

*/
