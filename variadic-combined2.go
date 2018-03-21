package main

import "fmt"
import "strings"

func main() {
	var hobbies = []string{"painting", "singing", "dancing"}
	yourHobbies("Qashwa", hobbies...)
}

func yourHobbies(name string, hobbies ...string) {
	var hobbiesAsString = strings.Join(hobbies, ", ")

	fmt.Printf("Hello, my name is: %s\n", name)
	fmt.Printf("My hobbies are: %s\n", hobbiesAsString)
}
