package main

import (
	"fmt"
	"os"
)

func main() {

	// divid by zero panic
	// for {
	// 	v := rand.IntN(10)
	// 	println(100 / (v - 1))
	// }

	// nil reference. panic: runtime error: invalid memory address or nil pointer dereference

	// ptr := new(int)

	// for {
	// 	v := rand.IntN(10)
	// 	if v == 0 {
	// 		ptr = nil // making the ptr as nill
	// 		*ptr = v  // dereference the value from nil
	// 	} else {
	// 		*ptr = v
	// 		println(*ptr)
	// 	}
	// }

	// func() { //main.func1
	// 	arr := [10]int{10, 11, 12, 13, 14, 15, 16, 17, 18, 10}
	// 	// panic: runtime error: index out of range [10] with length 10
	// 	for i := 0; i <= len(arr); i++ {
	// 		println(arr[i])
	// 	}
	// }()

	println("before panic")

	defer func() { // main.func3
		for i := 0; i <= 100; i++ {
			if i%10 == 0 {
				println("dib by 10", i)
			}
		}
	}()

	defer func() { // main.func3
		for i := 0; i <= 100; i++ {
			if i%9 == 0 {
				println("dib by 9", i)
			}
		}
	}()

	func(fn string) { //main.func1
		file, err := os.Open(fn)
		if err != nil {
			panic(err.Error())
		}

		defer file.Close()

		bytes := make([]byte, 100)

		n, err := file.Read(bytes)

		if err == nil {
			fmt.Println(string(bytes))
		} else {
			fmt.Println(err.Error())
		}

		if n == 0 {
			panic("no content in the file")
		}

	}("data.txt")

	func() { // main.func3
		for i := 0; i <= 100; i++ {
			if i%5 == 0 {
				println("dib by 5", i)
			}
		}
	}()
	//
}

// error --> can be handled
// panic --> by design, panic panics the whole call stack , nothing would executed after panic, but can recover from panic
// fatal --> os.Exit(n) , n!=0 any non zero is considered as crashed execution, cannot be recovered
