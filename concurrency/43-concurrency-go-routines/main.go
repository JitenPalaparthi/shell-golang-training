package main

import "time"

func main() {

	go func() {
		go println("Hello First Go routine")
	}()

	go func() { // main.func3
		for i := 0; i <= 100; i++ {
			if i%10 == 0 {
				time.Sleep(time.Nanosecond * 10)
				go println("dib by 10", i)
			}
		}
	}()

	go DivByNumber(4)

	println("End of main")
	time.Sleep(time.Millisecond * 10) // Just wait for 10 ms
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
