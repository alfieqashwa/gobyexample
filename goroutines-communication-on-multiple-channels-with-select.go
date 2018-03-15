/*
=== Communicating on Multiple Channels ===

PROBLEM
You would like to perform communication operations on multiple channels.

SOLUTION
Go provides a "select" statement that lets a goroutine perform communication operations on multiple channels.

HOW IT WORKS
When you build real-world concurrent programs with Go, you might need to deal with multiple channels in a single goroutine, which could require you to perform communication operations on multiple channels.
The "select" statement is a powerful communciation mechanism when it is used in conjunction with multiple channels. A "select" block is written with multiple case statements that a goroutine wait until one of the cases can run; it then executes the code block of that case.
If multiple case blocks are ready for execution, it randomly picks one of them and executes the code block of that case.

=== A SELECT BLOCK FOR READING VALUES FROM MULTIPLE CHANNELS ===
*/

package main

import (
	"fmt"
	"math"
	"math/rand"
	"sync"
	"time"
)

type (
	fibvalue struct {
		input, value int
	}
	squarevalue struct {
		input, value int
	}
)

func generateSquare(sqrs chan<- squarevalue) {
	defer wg.Done()
	for i := 1; i <= 10; i++ {
		num := rand.Intn(50)
		sqrs <- squarevalue{
			input: num,
			value: num * num,
		}
	}
}

func generateFibonacci(fibs chan<- fibvalue) {
	defer wg.Done()
	for i := 1; i <= 10; i++ {
		num := float64(rand.Intn(50))
		//Fibonacci using Binet's formula
		Phi := (1 + math.Sqrt(5)) / 2
		phi := (1 - math.Sqrt(5)) / 2
		result := (math.Pow(Phi, num) - math.Pow(phi, num)) / math.Sqrt(5)
		fibs <- fibvalue{
			input: int(num),
			value: int(result),
		}
	}
}

func printValues(fibs <-chan fibvalue, sqrs <-chan squarevalue) {
	defer wg.Done()
	for i := 1; i <= 20; i++ {
		select {
		case fib := <-fibs:
			fmt.Printf("Fibonacci value of %d is %d\n", fib.input, fib.value)
		case sqr := <-sqrs:
			fmt.Printf("Square values of %d is %d\n", sqr.input, sqr.value)
		case <-time.After(time.Second * 3): //  It also possible to implement a timeout expression
			fmt.Println("timed out") // in the select block
		}
	}
}

// wg is used to wait for the program to finish.
var wg sync.WaitGroup

func main() {

	wg.Add(3)
	// Create Channels
	fibs := make(chan fibvalue)
	sqrs := make(chan squarevalue)
	// Launching 3 goroutines
	go generateFibonacci(fibs)
	go generateSquare(sqrs)
	go printValues(fibs, sqrs)
	// Wait for completing all goroutines
	wg.Wait()
}
