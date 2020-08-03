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
	contactInfo
}

type personPointer *person

func (p person) print() {
	fmt.Printf("%+v\n", p)
}

func (pp *person) updateName(newFirstName string) {
	(*pp).firstName = newFirstName
}

func main() {
	alex := person{
		firstName: "Alex",
		lastName:  "Anderson",
		contactInfo: contactInfo{
			email:   "abc@test.com",
			zipCode: 123,
		},
		age: 20,
	}
	alex.updateName("Jimmy")
	alex.print()

	name := "bill"
	fmt.Println(*&name)
}
