package main

import (
	"fmt"
	"runtime"
	"sync"
	"time"
)

func main() {

	ch := make(chan string)
	sig := make(chan struct{})

	wg := new(sync.WaitGroup)
	wg.Add(3)

	go Generate(wg, "Gen-1", ch, time.Millisecond*1000)
	go Generate(wg, "Gen-2", ch, time.Millisecond*1000)
	go Generate(wg, "Gen-3", ch, time.Millisecond*1000)

	go func() {
		wg.Wait()
		close(ch)
	}()

	go func() {
		for r := range ch {
			println(r)
		}
		sig <- struct{}{}
	}()
	<-sig

}

func Generate(wg *sync.WaitGroup, name string, ch chan string, duration time.Duration) { // Sender chan
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
				//close(ch)
				wg.Done()
				runtime.Goexit()
				//default:
			}
		}
	}()
}

// v,ok:=<-ch
//
