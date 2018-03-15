/*
A Three-Stage Pipeline with Three Goroutines Connected by Two Channels

The example program has a three-stage pipeline with three goroutines that are connected by two channels. In this pipeline, the goroutine of the first stage is used to randomly generate values with an upper limit of 50. The pipeline has an outbound channel to give inbound values to the goroutine of the second stage.
The goroutinee of the second stage has an inbound channel and an outbound channel. The inbound channel receives values from the first goroutine when it randomly generates each value and finds out the Fibonacci value. It then provides the resulting Fibonacci values to the goroutine of third stage, which just prints the outbound values from the goroutine of second stage.
Here is the example program.
*/

package main

import (
	"fmt"
	"math"
	"math/rand"
	"sync"
)

type fibvalue struct {
	input, value int
}

var wg sync.WaitGroup

// Generates random values
func randomCounter(out chan int) {
	defer wg.Done()
	var random int
	for x := 0; x < 10; x++ {
		random = rand.Intn(50)
		out <- random
	}
	close(out)
}

// Produces Fibonacci values of inputs provided by randomCounter
func generateFibonacci(out chan fibvalue, in chan int) {
	defer wg.Done()
	var input float64
	for v := range in {
		input = float64(v)
		// Fibonacci using Binet's formula
		Phi := (1 + math.Sqrt(5)) / 2
		phi := (1 - math.Sqrt(5)) / 2
		result := (math.Pow(Phi, input) - math.Pow(phi, input)) / math.Sqrt(5)
		out <- fibvalue{
			input: v,
			value: int(result),
		}
	}
	close(out)
}

// Print Fibonacci values generated bt generateFibonacci
func printFibonacci(in chan fibvalue) {
	defer wg.Done()
	for v := range in {
		fmt.Printf("Fibonacci value of %d is %d\n", v.input, v.value)
	}
}

func main() {
	// Add 3 into WaitGroup Counter
	wg.Add(3)
	// Declare Channels
	randoms := make(chan int)
	fibs := make(chan fibvalue)
	// Launching 3 goroutines
	go randomCounter(randoms)           // First stage of pipeline
	go generateFibonacci(fibs, randoms) // Second stage of pipeline
	go printFibonacci(fibs)             // Third stage of pipeline
	// Wait for completing all goroutines
	wg.Wait()
}
