package main

import (
	"sync"
	"time"
)

func main() {

	wg := new(sync.WaitGroup)
	ch := make(chan int)
	wg.Add(2)
	go publish(wg, 10, ch)
	go consume(wg, 10, ch)
	wg.Wait()
}

// func publish(wg *sync.WaitGroup, r uint, ch chan int) {

func publish(wg *sync.WaitGroup, r uint, ch chan<- int) { // Sender chan
	a, b := 0, 1
	for i := uint(1); i <= r; i++ {
		time.Sleep(time.Millisecond * 100)
		ch <- a
		a, b = b, a+b
	}
	wg.Done()
}

//func consume(wg *sync.WaitGroup, r uint, ch chan int) {

func consume(wg *sync.WaitGroup, r uint, ch <-chan int) { // receiver ch
	for i := uint(1); i <= r; i++ {
		println(<-ch)
	}
	wg.Done()
}
