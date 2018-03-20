package main

import "fmt"

func main() {

	var chicken = map[string]int{
		"jan":   50,
		"feb":   40,
		"march": 30,
		"april": 20,
	}

	for key, val := range chicken {
		fmt.Println(key, "\t", val)
	}
}
