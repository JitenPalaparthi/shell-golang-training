package main

import "fmt"

func main() {

	c1 := Customer{Id: 100, Name: "Jiten", Email: "jitenP@outlook.com", Status: "active", Address: Address{City: "Blr", Pincode: "560086", Status: "current"}}
	println(c1.Address.City) // either can access like this
	println(c1.City)         // or like this directly
	println(c1.Pincode)
	println("Customer Status", c1.Status)
	println("Address Status", c1.Address.Status)

	cc1 := ColourCode{int: 100, string: "Full Red", ShortName: "red"}

	fmt.Println(cc1)

	fmt.Println("id", cc1.int)
	fmt.Println("code", cc1.string)

	cc1.ShortName = "Red" // mutate directly with type
	fmt.Println("short code", cc1.ShortName)

}

type Customer struct {
	Id      int
	Name    string
	Email   string
	Status  string
	Address // Promoted field ,create a file without a field name, just a type name
}

type Address struct {
	City, Pincode, Status string
}

type ColourCode struct { // a struct with anonymous fields
	int       // no name to the field only the type
	string    // no name to the field only the type
	ShortName // if string is to be repeated , can create an alias and use it here
}

type ShortName = string // alias
