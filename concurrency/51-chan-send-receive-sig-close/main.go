package main

import (
	"runtime"
	"time"
)

func main() {

	ch := make(chan int)

	sig := make(chan struct{})

	go publish(10, ch)
	go consume("consumer-1", ch, sig)
	go consume("consumer-2", ch, sig)

	<-sig
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

// close does not make the chanel nil
// once the sender closes a channel cannot open it or something
// only the receiver can know that the channel is closed or not
