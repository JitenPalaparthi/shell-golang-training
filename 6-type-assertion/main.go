package main

import (
	"fmt"
	"strconv"
)

// package main is create and bootstraped by go runtime
// no need to install any runtime, it is embedded with every binary of Go

func main() {

	var any1 any // interface

	if any1 == nil {
		println("yes it is nil")
	}

	any1 = uint8(123)

	// Data Ptr --> Pointer to the value 123
	// Type Ptr --> Pointer to the type called uint8

	var num1 uint8 = any1.(uint8) // cannot do type casting , assertion
	// any.(T)
	println(num1)
	var num2 int

	// num2 = any1.(int)
	// println(num2)

	num2, ok := any1.(int) // this type assertion return one or two values
	if ok {
		println(num2)
	} else {
		println("could not be successfully asserted")
	}

	println("Hello World")

	var (
		n1       = uint8(123)
		n2       = 123123
		n3       = float32(12.23)
		n4       = 82323.23
		n5       = 0b11001100110011 // int in the form of binary
		n6       = 0xaabbccddeeff   // int in the form of hex
		n7       = uint64(123434343)
		any2     = any("1231231")
		any3     = any(true)
		any4 any = 2131231
		any5 any = 123123.123
		str1     = "34234234"
		str2     = "123213.21312"
	)

	var sum float64 = float64(n1) + float64(n2) + float64(n3) + n4 + float64(n7) + float64(n5) + float64(n6)

	sum1, ok := any4.(int)
	if ok {
		sum += float64(sum1)
	}

	sum2, ok := any5.(float64)

	if ok {
		sum += sum2
	}

	str3, ok := any2.(string)

	if ok {
		v, err := strconv.ParseFloat(str3, 64)
		if err == nil {
			sum += v
		}
	}

	b1, ok := any3.(bool)

	if ok {
		if b1 {
			sum += 1 // float64 can be aqdded to 1 , go automatically uses 1.0
		}
	}

	v, err := strconv.ParseInt(str1, 2, 64)
	if err == nil {
		sum += float64(v)
	}

	v1, err := strconv.ParseFloat(str2, 64)
	if err == nil {
		sum += v1
	}

	fmt.Printf("Sum:%.3f", sum)

}

// type cast
// conversion
// assertion

type Shape interface {
	Area() float32
}

type Empty interface{} // everybody implements this
