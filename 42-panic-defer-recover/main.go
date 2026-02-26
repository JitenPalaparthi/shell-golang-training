package main

import (
	"fmt"
	"os"
)

func main() {

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
		defer func() {
			if r := recover(); r != nil {
				fmt.Println("recover me..", r)
			}
		}()

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

func Fatal(msg string) {
	println(msg)
	os.Exit(1)
}

// error --> can be handled
// panic --> by design, panic panics the whole call stack , nothing would executed after panic, but can recover from panic
// fatal --> os.Exit(n) , n!=0 any non zero is considered as crashed execution, cannot be recovered
