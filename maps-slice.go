package main

import "fmt"

func main() {
	var chickens = []map[string]string{

		{"name": "chicken blue", "gender": "male"},
		{"name": "chicken red", "gender": "male"},
		{"name": "chicken yellow", "gender": "female"},
		// OR
		// map[string]string{"name": "chicken blue", "gender": "male"},
		// map[string]string{"name": "chicken red", "gender": "male"},
		// map[string]string{"name": "chicken yellow", "gender": "female"},

	}

	for _, chicken := range chickens {
		fmt.Println(chicken["gender"], chicken["name"])
	}
}
