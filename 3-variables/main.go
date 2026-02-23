package main

import (
	"fmt"
	"math/rand/v2"
	"reflect" // used a lot while using tags
)

// package main is create and bootstraped by go runtime
// no need to install any runtime, it is embedded with every binary of Go

const ( // RO data segment
	TICK   = 60
	SECOND = 1
	MIN    = SECOND * TICK
	HOUR   = MIN * TICK
	SOME   = TICK + 10*(SECOND*5) + HOUR/30 // when does --> by the compiler at compiler
)

var ( // Global variable , static variable
	rnd        = rand.IntN(100) //GetValue()
	Global int = 1000           // never be deallocated similar to static variables
)

// constants are stored in RO data segment
// can I evaluate constant --> Yes if the operands are constatns
// When does it evaluage --> by the compiler at compiletime

func main() {

	const PI = 3.14 // type infafence
	const PI1 float32 = 3.14

	// var a = 100                                      // would be assigned at runtime
	//const SOME1 = TICK + 10*(SECOND*5) + HOUR/30 + a // some silly stuff

	var (
		//tick   = 60
		second = 1
		hour   = TICK * TICK * second
	)

	some := TICK + 10*(second*5) + hour/30 // some silly stuff
	println(some)

	println(TICK, HOUR, PI)

	fmt.Println("Type of PI:", reflect.TypeOf(PI))
	Global++

	// 100178 4a8 D main.Global
	// 102f04 4a8 // data segment
	// 4d2721108020 // stack
	fmt.Printf("%p\n", &Global)
	fmt.Printf("%p\n", &some)
}

func GetValue() int { // inline code
	// a, b := 10, 20
	// println(a, b)
	// a = a * a
	// b = b * b
	// c := a + b
	return rand.IntN(100)
}

func greet() {
	println("Hello World")
}
