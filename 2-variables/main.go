package main

// package main is create and bootstraped by go runtime
// no need to install any runtime, it is embedded with every binary of Go

import (
	"fmt"
)

func main() {

	var num1 int = -1000 // allocated on stack

	var num2 uint = 123123

	var num3 uint8 = 255

	var float1 float32 = 12312.123

	var float2 float64 = 312.34434343

	var (
		num4, num5 float32 // = 10.6, 20.5 // what are the values of this
	)

	var ok1 = true // type inference , the type is bool

	var num6 = 1234123 // int by type inference

	var float3 = 9324324.344343 // float64

	var float4 = 1.34 // float64

	var age = 41 // int --> 8 bytes on 64 bit machine

	var age1 uint8 = 41

	var str string // what is the value of this
	var ok2 bool   // what is the valuer of this?

	fmt.Println(num1, num2, num3, num4, num5, num6, float1, float2, float3, float4, ok1, ok2, age, age1, str)
	fmt.Printf("%.2f\n", float3)
	fmt.Printf("age:%d type of age: %T", age, age)

	// shorthand declaration

	ok3 := true // bool

	str1 := "Hello Shell" // type inferance --> string

	a, b := 10, 20 // Type inferance -->int

	// swap

	t := a // t is a variable
	a = b
	b = t

	a, b = b, a // simply do this , it does the swapping at a assembly level, it would ahve fewer instructions than the temp variable swap

	a1, b1, c1 := 10, 20, 30

	a1, b1, c1 = b1, c1, a1 // swap more than two

	println(ok3, str1)
	println(a, b)
	println(a1, b1, c1)

}

// Numbers
// uint,int, uint8,uint16,uint32,uint64, int8,int16,int32,int64,flaot32,float64
// byte | uint8, rune | int32
// complex,uintptr

// boolean --> bool

// strings --> string

// anything --> any | interface{}(<1.18 there is empty interface, created an alias called any)

// Type inference
// Zero value , based type, for numbers it is 0, for bool -> false, for string ->""

// Go is statically typed langauge
