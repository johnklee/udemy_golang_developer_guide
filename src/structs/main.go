package main

import "fmt"

type contactInfo struct {
	email   string
	zipCode int
}

type person struct {
	firstName string
	lastName  string
	age       int
	contact   contactInfo
}

func main() {
	alex := person{
		firstName: "Alex",
		lastName:  "Anderson",
		contact: contactInfo{
			email:   "abc@test.com",
			zipCode: 123,
		},
		age: 20,
	}

	fmt.Printf("%+v\n", alex)
}
