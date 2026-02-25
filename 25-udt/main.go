package main

import "fmt"

func main() {

	var c1 Customer // c1 can be on stack or on heap based on
	c1.Id = 100
	c1.Name = "Jiten"
	c1.Email = "JitenP@Outlook.com"
	c1.Status = "active"
	fmt.Println(c1)

	e1 := Employee{Id: 101, Name: "Jiten", Email: "JitenP@outlook.com", Mobile: "9618558500", Status: "active", Address: &Address{City: "Bangalore", Pincode: "560086", Status: "active"}}
	fmt.Println(e1)
	fmt.Println(*(e1.Address))
	fmt.Println("City", e1.Address.City)

	e2 := &Employee{Id: 101, Name: "Jiten", Email: "JitenP@outlook.com", Mobile: "9618558500", Status: "active", Address: &Address{City: "Bangalore", Pincode: "560086", Status: "active"}}
	fmt.Println(e2)
	fmt.Println("City", (*e2).Address.City)
	fmt.Println("Email", e2.Email)
	fmt.Println("Mobile", (*e2).Email)

	e3 := new(Employee)
	(*e3).Id = 101
	e3.Email = "Jitenp@outlook.com"
	e3.Mobile = "9618558500"
	// e3.Address.City = "Bangalore"
	// e3.Address.Pincode = "560086"
	// e3.Address.Status = "current"

	addr1 := &Address{}
	addr1.City = "Bangalore"
	addr1.Pincode = "560086"
	addr1.Status = "current"
	e3.Address = addr1

	fmt.Println(e3)

	fmt.Println(*(e3.Address)) // no need to dereference

	addr2 := e3.Address

	fmt.Println(addr2)
	fmt.Println(addr2.City)

	fmt.Println((*addr2).City)

	e3.Address = nil
	fmt.Println(e3)

}

type Customer struct {
	Id     int
	Name   string
	Email  string
	Status string
}

type Employee struct {
	Id                          int
	Name, Email, Mobile, Status string
	Address                     *Address // can be same, it is a composition
}

type Address struct {
	City, Pincode, Status string
}
