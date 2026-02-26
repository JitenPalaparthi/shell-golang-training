package main

import (
	"fmt"
	"runtime"
	"sync"
	"time"
)

func main() {

	ch1 := Generate("chan-1", time.Second*1)
	ch2 := Generate("chan-2", time.Millisecond*1200)
	ch3 := Generate("chan-3", time.Millisecond*1100)
	out := make(chan string, 10)

	go FanIn(out, ch1, ch2, ch3)
	sig := make(chan struct{})
	go func() {
		for c := range out {
			println(c)
		}
		sig <- struct{}{}
	}()
	<-sig

}

func Generate(name string, duration time.Duration) chan string { // Sender chan
	ch := make(chan string, 5)
	//chTimout := TimeOut(duration)
	chTimout := time.After(duration)
	go func() {
		i := 1
		for {
			select { // only works with channels can be senders, receivers
			case ch <- fmt.Sprint("Generator", name, "-->", i*i):
				i++
			case <-chTimout:
				//fmt.Println("Time---->", v)
				close(ch)
				runtime.Goexit()
				//default:
			}
		}
	}()
	return ch
}

func FanIn(out chan<- string, chs ...chan string) {
	wg := new(sync.WaitGroup)
	wg.Add(len(chs))

	for _, ch := range chs {
		go func() {
			for c := range ch {
				//println(c)
				// do some processing
				out <- fmt.Sprint("Sending to Out channel-->", c)
			}
			wg.Done()
		}()
	}

	wg.Wait()
	close(out)

}
