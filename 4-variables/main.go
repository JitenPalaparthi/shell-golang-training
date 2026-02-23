package main

import (
	"fmt"
	"unsafe"
)

// package main is create and bootstraped by go runtime
// no need to install any runtime, it is embedded with every binary of Go

func main() {

	str1 := "Hello 世界 😀" // 8 chars, 17 bytes

	str1 = "Hello Shell"

	str1 = "Hello World!"

	var str2 string // This is not nil
	//var str3, str4 string

	println(len(str1), "size of str1:", unsafe.Sizeof(str1), "str2:", unsafe.Sizeof(str2))

	// reflect.TypeOf()
	// unsafe.SizeOf()

	// any datatype

	var any1 any

	// What is any

	// can assign any kind of data to any variable

	if any1 == nil {
		println("yes it is nil")
	}

	fmt.Printf("data ptr:%v type ptr:%T", any1, any1)

	any1 = true

	any1 = "Hello Wrold"

	any1 = 23.23

	any1 = 0b11001100

	any1 = 0x1234faaff

	any1 = str2

	fmt.Println("size of any1:", unsafe.Sizeof(any1))

	//any1 = Person{Id: 101}

}

// type Person struct {
// 	Id   int
// 	name string
// 	age  uint8
// }

// 32 bit --> 4gb
// 8 gb --> 4gb , 3500rs
// 32bit --> 4 bytes
