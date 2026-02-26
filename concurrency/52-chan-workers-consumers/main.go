package main

import (
	"runtime"
	"sync"
	"time"
)

func main() {

	ch := make(chan int, 5)

	sig := make(chan struct{})

	go publish(10, ch)
	// go consume("consumer-1", ch, sig)
	// go consume("consumer-2", ch, sig)
	// go receiver("consumer-3", ch, sig)

	// <-sig
	// <-sig
	// <-sig

	go worker(5, ch, sig)

	<-sig

}

func publish(r uint, ch chan<- int) { // Sender chan
	a, b := 0, 1
	for i := uint(1); i <= r; i++ {

		ch <- a
		a, b = b, a+b
	}
	close(ch) // close closes a channel , senders side

}

func consume(name string, ch <-chan int, sig chan<- struct{}) { // receiver ch
	for {
		time.Sleep(time.Millisecond * 10)
		v, ok := <-ch
		if ok {
			println("consuming by-->", name, "--> ", v)
		} else {
			//break
			sig <- struct{}{}
			//close(sig)
			runtime.Goexit()
		}
	}
}

func receiver(name string, ch <-chan int, sig chan<- struct{}) {
	for v := range ch { // range loop iterate until the channel is closed
		time.Sleep(time.Millisecond * 10)
		println("consuming by-->", name, "--> ", v)
	}
	sig <- struct{}{}
	runtime.Goexit()
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
