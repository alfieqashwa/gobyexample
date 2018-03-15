package main

import "fmt"

type Address struct {
	Street, City, State, Zip string
	IsShippingAddress        bool
}

type Customer struct {
	FirstName, LastName, Email, Phone string
	Addresses                         []Address
}

func (c Customer) ToString() string {
	return fmt.Sprintf("Customer: %s %s, Email: %s", c.FirstName, c.LastName, c.Email)
}
func (c *Customer) ChangeEmail(newEmail string) { // *Customer is Pointer Receiver
	c.Email = newEmail
}

func (c Customer) ShippingAddress() string {
	for _, v := range c.Addresses {
		if v.IsShippingAddress == true {
			return fmt.Sprintf("%s, %s, %s, Zip: %s", v.Street, v.City, v.State, v.Zip)
		}
	}
	return ""
}

func main() {
	c := Customer{
		FirstName: "Alfie",
		LastName:  "Qashwa",
		Email:     "alfieqashwa@gmail.com",
		Phone:     "0812-8093-19xx",
		Addresses: []Address{
			Address{
				Street:            "Komp Diknas",
				City:              "Tangsel",
				State:             "ID",
				Zip:               "15xxx",
				IsShippingAddress: true,
			},
			Address{
				Street: "Nerada TownHouse 1",
				City:   "Tangsel",
				State:  "ID",
				Zip:    "15xxx",
			},
		},
	}
	fmt.Println(c.ToString())
	fmt.Println(c.ShippingAddress())
	// Call ChangeEmail
	c.ChangeEmail("circleq.co@gmail.com")
	fmt.Println(c.ToString())
}

/*
=== OUTPUT ===

Customer: Alfie Qashwa, Email: alfieqashwa@gmail.com
Komp Diknas, Tangsel, ID, Zip: 15xxx
Customer: Alfie Qashwa, Email: circleq.co@gmail.com

*/
