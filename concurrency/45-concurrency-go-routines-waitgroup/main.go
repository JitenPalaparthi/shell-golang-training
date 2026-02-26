package main

import (
	"runtime"
	"sync"
	"time"
)

var wg = new(sync.WaitGroup)

func main() {

	defer wg.Wait() // it waits until the state is zero

	wg.Add(1)
	go func() {
		defer wg.Done()
		println("Hello First Go routine")

	}()

	wg.Add(1)
	go func() {
		wg.Done()
	}()

	wg.Add(1)
	go func() { // main.func3
		i := 1
		for {

			if i > 100 {
				wg.Done()
				runtime.Goexit()
			}

			if i%10 == 0 {
				time.Sleep(time.Nanosecond * 10)
				println("dib by 10", i)
			}
			i++
		}
	}()

	wg.Add(1)
	go func() {
		DivByNumber(4)
		wg.Done()
	}()
	wg.Add(1)
	go DivByNumber2(wg, 5)

	wg.Add(1)
	go func(wg *sync.WaitGroup) {
		DivByNumber(6)
		wg.Done()
	}(wg)

	println("End of main")
	// go func() {
	// 	for {
	// 	}
	// }()

	//'runtime.Goexit() // failure exit

	//wg.Wait()
}

func DivByNumber(x int) {
	for i := 0; i <= 100; i++ {
		if i%x == 0 {
			time.Sleep(time.Nanosecond * 10)
			go println("div by ", x, "=", i)
		}
	}
}

func DivByNumber2(wg *sync.WaitGroup, x int) {
	for i := 0; i <= 100; i++ {
		if i%x == 0 {
			time.Sleep(time.Nanosecond * 10)
			go println("div by ", x, "=", i)
		}
	}
	wg.Done()
}

// 1. main is also a goroutine
// 2. goroutines cannot run in a order
// 3. main does not wait for other goroutines to complete its execute, no goroutine waits for other to complete its execution by default
