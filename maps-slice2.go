package main

import "fmt"

func main() {
	var chickens = []map[string]string{

		{"name": "chicken blue", "gender": "male"},
		{"address": "mangga street", "id": "k001"},
		{"community": "chicken lovers"},
	}

	for _, chicken := range chickens {
		fmt.Println(chicken["gender"], chicken["name"])
		fmt.Println(chicken["address"], chicken["id"])
		fmt.Println(chicken["community"])
	}
}
