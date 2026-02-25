package main

import (
	"fmt"
	"unsafe"
)

func main() {
	c1 := Customer{Id: 100, Name: "Jiten", Email: "jitenP@outlook.com", Status: "active", Address: struct {
		City    string
		Pincode string
		Status  string
	}{City: "Blr", Pincode: "560086", Status: "current"}}
	println(c1.Address.City) // either can access like this
	println(c1.Address.City) // or like this directly
	println(c1.Address.Pincode)
	println("Customer Status", c1.Status)
	println("Address Status", c1.Address.Status)

	var addr struct { // anonymous struct as a variable
		City, Pincode, Status string
	}

	addr.City = "hyd"
	addr.Pincode = "54000"
	addr.Status = "active"

	var fn func() = Greet
	if fn != nil {
		fn()
	} else {
		println("it is nil func")
	}
	// func is a poiter

	fmt.Printf("%p\n", &c1)

	uptr := uintptr(unsafe.Pointer(&c1))
	ptr := (*int)(unsafe.Pointer(uptr)) //4, 1, 2, 3
	//ptr2 := (*int)(unsafe.Pointer(uintptr(unsafe.Pointer(&c1))))
	fmt.Println(*ptr)

	fmt.Printf("%p\n", &Empty1{}) // runtime.zerobase
	fmt.Printf("%p\n", &Empty2{}) // runtime.zerobase
	// empoty slice  without nil

	arr := [0]int{}
	fmt.Printf("%p\n", &arr) // runtime.zerobase

}

func Greet() {
	println("hello Shell minds!")
}

type Customer struct {
	Id      int      //8
	Name    string   //16
	Email   string   //16
	Status  string   //16
	Address struct { // embedded struct
		City, Pincode, Status string //16 //16 //16
	}
}

func PrintAddress(a struct { // anonymous struct as a variable
	City, Pincode, Status string
}) {
	fmt.Println("City:", a.City)
	fmt.Println("PinCode:", a.Pincode)
	fmt.Println("Status:", a.Status)
}

type Empty1 struct{}

type Empty2 struct{}
