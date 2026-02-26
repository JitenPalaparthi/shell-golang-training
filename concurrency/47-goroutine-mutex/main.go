package main

import "sync"

var (
	mu      = new(sync.Mutex)
	wg      = new(sync.WaitGroup)
	Counter = 1
)

func main() {

	wg := new(sync.WaitGroup)
	wg.Add(2)

	go func() {
		for i := 1; i <= 1000; i++ {
			mu.Lock()
			Counter++
			mu.Unlock()
		}
		wg.Done()
	}()

	go func() {
		for i := 1; i <= 1000; i++ {
			mu.Lock()
			Counter--
			mu.Unlock()
		}
		wg.Done()
	}()

	wg.Wait()
	println(Counter)

}
