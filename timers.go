package main

import (
	"fmt"
	"time"
)

/*
Timers represent a single event in the future.
You tell the timer how long you want to wait,
and it provides a channel that will be notified at that time.
This timer will wait 2 seconds.
*/
func main() {
	timer1 := time.NewTimer(2 * time.Second)

	<-timer1.C
	fmt.Println("Timer 1 expired")

	timer2 := time.NewTimer(time.Second)
	go func() {
		<-timer2.C
		fmt.Println("Timer 2 Expired")
	}()
	stop2 := timer2.Stop()
	if stop2 {
		fmt.Println("Timer 2 stopped")
	}
}
