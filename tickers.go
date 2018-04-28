/*
Timers are for when you want to do something once in the future -
tickers are for when you want to do something repeatedly at regular intervals.
Here’s an example of a ticker that ticks periodically until we stop it
*/
package main

import (
	"fmt"
	"time"
)

func main() {
	ticker := time.NewTicker(500 * time.Millisecond)
	go func() {
		for t := range ticker.C {
			fmt.Println("Tick at", t)
		}
	}()
	// Tickers can be stopped like timers.
	// Once a ticker is stopped it won't receive any more values on its channel.
	// We'll stop ours after 5500ms.
	time.Sleep(5500 * time.Millisecond)
	ticker.Stop()
	fmt.Println("Ticker stopped")
}
