package main

import (
	"runtime"
	"time"
)

func main() {

	go func() {
		go println("Hello First Go routine")
	}()

	go func() {

	}()

	go func() { // main.func3
		i := 1
		for {

			if i > 10 {
				runtime.Goexit()
			}

			if i%10 == 0 {
				time.Sleep(time.Nanosecond * 10)
				go println("dib by 10", i)
			}
			i++
		}
	}()

	go DivByNumber(4)

	println("End of main")
	// go func() {
	// 	for {
	// 	}
	// }()

	runtime.Goexit() // failure exit
}

func DivByNumber(x int) {
	for i := 0; i <= 100; i++ {
		if i%x == 0 {
			time.Sleep(time.Nanosecond * 10)
			go println("div by ", x, "=", i)
		}
	}
}

// 1. main is also a goroutine
// 2. goroutines cannot run in a order
// 3. main does not wait for other goroutines to complete its execute, no goroutine waits for other to complete its execution by default
