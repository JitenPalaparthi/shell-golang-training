package main

import (
	"runtime"
	"sync"
	"time"
)

func main() {

	sig := make(chan struct{})

	ch1 := Generate(time.Second * 2)

	go worker(5, ch1, sig)

	<-sig

}

func Generate(duration time.Duration) chan int { // Sender chan
	ch := make(chan int, 5)
	//chTimout := TimeOut(duration)
	chTimout := time.After(duration)
	go func() {
		i := 1
		for {
			select { // only works with channels can be senders, receivers
			case ch <- i * i:
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

func TimeOut(duration time.Duration) chan struct{} {
	ch := make(chan struct{})
	go func() {
		time.Sleep(duration)
		ch <- struct{}{}
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
