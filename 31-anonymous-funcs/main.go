package main

import "fmt"

func main() {

	func() {
		println("Hello World")
	}() // Anonymous function

	var fn1 func()
	fn1 = func() {
		println("Hello World")
	}

	if fn1 != nil {
		fn1()
	}

	var fn2 func(int, int) int

	fn2 = func(i1, i2 int) int {
		return i1 + i2
	}

	r1 := fn2(10, 20)

	fmt.Println("sum:", r1)

	fibs := func(r uint) []int {
		var fibs []int
		a, b := 0, 1
		for i := uint(1); i <= r; i++ {
			fibs = append(fibs, a)
			a, b = b, a+b
		}
		return fibs
	}(10)
	fmt.Println(fibs)

	var fn Fn = func() {
		println("Creating an anonymous function")
	}

	fn()       // can ezecute as a func
	fn.Greet() // can call method by using fn as a variable of a type
	fn.Hey()()
}

type Fn func()

func (f Fn) Greet() {
	println("hello Fn from Greet")
}

func (f Fn) Hey() Fn {
	println("calling hey and then mutating the orignal fn object")
	sum := 0
	f = func() {
		println("Mutating the func")
		for i := 1; i <= 10; i++ {
			sum += i
		}
		println("Sum:", sum)
	}

	return f
}
