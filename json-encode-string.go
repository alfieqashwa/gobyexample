// This sample program demonstrates how to encode a JSON string.
package main

import (
	"encoding/json"
	"fmt"
)

func main() {
	// Create a map of key/value pairs and parses the data into JSON
	emp := make(map[string]interface{})
	emp["name"] = "Alfie Qashwa"
	emp["jobtitle"] = "Software Developer"
	emp["phone"] = map[string]interface{}{
		"home":   "021-749-0xxx",
		"office": "0812-8093-1xxx",
	}
	emp["email"] = "alfieqashwa@gmail.com"

	// Unmarshal the map into a JSON string.
	empData, err := json.Marshal(emp)
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	jsonStr := string(empData)
	fmt.Println("The JSON data is:")
	fmt.Println(jsonStr)
}

/*
=== OUTPUT ===
The JSON data is:
{"email":"alfieqashwa@gmail.com","jobtitle":"Software Developer","name":"Alfie Qashwa","phone":{"home":"021-749-0xxx","office":"0812-8093-1xxx"}}

=== SYNTAX ===
func Marshal(v interfaace{}) ([]byte, error)

=== EXPLANATION ===
Marshal function is good for producing JSON that could be returned in a network response,
like a Restfult API. The function Marshal returns two values:
the encoded JSON data as slice byte and ane error value.
Using Marshal, we can also encode the values of struct type as json values that will help us
to quickly build JSON-based APIs.

*/
