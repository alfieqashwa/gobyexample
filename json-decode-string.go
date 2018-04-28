// This sample program demonstrates how to decode a JSON string.
package main

import (
	"encoding/json" // Encoding and decoding Package
	"fmt"
)

// JSON Contains a sample String to unmarshal.
var JSON = `{
	"name":"Mark Taylor",
	"jobtitle":"Software Developer",
	"phone":{
		"home":"123-456-887",
		"office":"234-658-545"
	},
	"email":"markt@gmailcom"
}`

func main() {
	// Unmarshall the JSON string into info map variable.
	var info map[string]interface{}
	json.Unmarshal([]byte(JSON), &info)

	// Print the output from info map.
	fmt.Println(info["name"])
	fmt.Println(info["jobtitle"])
	fmt.Println(info["email"])
	fmt.Println(info["phone"].(map[string]interface{})["home"])
	fmt.Println(info["phone"].(map[string]interface{})["office"])
}

/*
Mark Taylor
Software Developer
markt@gmailcom
123-456-887
234-658-545
*/
