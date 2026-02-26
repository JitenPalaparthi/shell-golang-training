package main

import "time"

// Do not communicate by sharing memory; instead, share memory by communicating.
// channels

func main() {
	// var ch chan int // this is declaration
	// if ch == nil {
	// 	println("chan is nil")
	// } else {
	// 	ch = make(chan int) // unbuffered
	// 	//ch = make(chan int, 10) // buffered
	// }
	ch := make(chan int, 2) // make is for slice, map and chan

	// to send a value to the channel
	go func() {
		//time.Sleep(time.Second * 2)
		ch <- 100 // ch <-
		println("Value has been sent-1")
		ch <- 101
		println("Value has been sent-2")
		ch <- 102
		println("Value has been sent-3")
		ch <- 104
		println("Value has been sen-4")
	}()
	time.Sleep(time.Second * 4)
	v := <-ch // The main is blocked here to receive a value .. waitin to receive a value from another go routine
	println(v)
	v = <-ch // The main is blocked here to receive a value .. waitin to receive a value from another go routine
	println(v)
	v = <-ch // The main is blocked here to receive a value .. waitin to receive a value from another go routine
	println(v)
	v = <-ch // The main is blocked here to receive a value .. waitin to receive a value from another go routine
	println(v)
}

// channel --> chan
// Sender
// receiver
// channel (pipe or querue)

// unbuffered
// buffered

// on umbuffered channel , the sender is blocked until the receiver receives the data
// on umbuffered channel, the receiver is blocked until the serder sends the data

// if there is no communication, both the sender and receiber are blocked
// On buffered --> The sender is not blocked until the buffer is full
