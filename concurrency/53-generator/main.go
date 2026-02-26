package main

import (
	"sync"
	"time"
)

func main() {

	sig := make(chan struct{})

	ch1 := Generate(10)

	go worker(5, ch1, sig)

	<-sig

}

func Generate(r uint) chan int { // Sender chan
	ch := make(chan int, 5)
	go func() {
		a, b := 0, 1
		for i := uint(1); i <= r; i++ {
			ch <- a
			a, b = b, a+b
		}
		close(ch) // close closes a channel , senders side
	}()
	return ch
}

func worker(num uint8, ch <-chan int, sig chan<- struct{}) {
	wg := new(sync.WaitGroup)
	wg.Add(int(num) + 1)

	go func() {
		wg.Wait()
		sig <- struct{}{}
		close(sig)
	}()

	go func() {
		for i := range num {
			go func() {
				for v := range ch { // range loop iterate until the channel is closed
					time.Sleep(time.Millisecond * 10)
					println("consuming by-->", i+1, "--> ", v)
				}
				wg.Done()
			}()
		}
		wg.Done()
	}()
}
