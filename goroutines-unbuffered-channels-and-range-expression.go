package main

import (
	"fmt"
	"sync"
)

// wg is used to wait tor the program to finish
var wg sync.WaitGroup

func main() {

	count := make(chan int)
	// Add a count of two, one for each goroutine.
	wg.Add(2)

	fmt.Println("Start Goroutines")
	// Launch a goroutine with label "Goroutine-1"
	go printCounts("Goroutine-1", count)
	// Launch a goroutine with label "Goroutine-2"
	go printCounts("Goroutine-2", count)
	fmt.Println("Communicaion of channel begins")

	count <- 1
	// Wait for the goroutines to finish.
	fmt.Println("Waiting To Finish")
	wg.Wait()
	fmt.Println("\nTerminating the Program")
}

func printCounts(label string, count chan int) {
	// Schedule the call to WaitGroup's Done to tell goroutine is completed.
	defer wg.Done()
	for val := range count {
		fmt.Printf("Count: %d received from %s \n", val, label)
		if val == 10 {
			fmt.Printf("Channel Closed from %s \n", label)
			// Close the channel
			close(count)
			return
		}
		val++
		// Send count back to the other goroutine.
		count <- val
	}
}

/*
=== OUTPUT ===

Waiting To Finish
Count: 1 received from Goroutine-1
Count: 2 received from Goroutine-2
Count: 3 received from Goroutine-1
Count: 4 received from Goroutine-2
Count: 5 received from Goroutine-1
Count: 6 received from Goroutine-2
Count: 7 received from Goroutine-1
Count: 8 received from Goroutine-2
Count: 9 received from Goroutine-1
Count: 10 received from Goroutine-2
Channel Closed from Goroutine-2

Terminating the Program

*/
