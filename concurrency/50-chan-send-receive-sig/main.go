package main

import (
	"time"
)

func main() {

	ch := make(chan int)
	ch1 := make(chan int)
	//sig := make(chan bool)
	//sig := make(chan Empty)
	sig := make(chan struct{})

	go publish(10, ch)
	go publish(10, ch1)
	go consume(10, ch, sig)

	for i := 1; i <= 10; i++ {
		println(<-ch1)
	}
	<-sig

	//ok := <-sig
	//println(ok)
	// a := 10
	// _ = a
}

func publish(r uint, ch chan<- int) { // Sender chan
	a, b := 0, 1
	for i := uint(1); i <= r; i++ {
		time.Sleep(time.Millisecond * 100)
		ch <- a
		a, b = b, a+b
	}

}

func consume(r uint, ch <-chan int, sig chan<- struct{}) { // receiver ch

	//func consume(r uint, ch <-chan int, sig chan<- Empty) { // receiver ch

	//func consume(r uint, ch <-chan int, sig chan<- bool) { // receiver ch
	for i := uint(1); i <= r; i++ {
		println(<-ch)
	}

	//sig <- true
	// sig <- Empty{}
	sig <- struct{}{}
}

type Empty struct{}
