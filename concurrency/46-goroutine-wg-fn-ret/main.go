package main

import "sync"

var Sum1 int

var Sum2 int

func main() {
	wg := new(sync.WaitGroup)

	wg.Add(1)
	go func() {
		Sum1 = SumOfFibNumbers(10)
		wg.Done()
	}()

	wg.Add(1)
	go func() {
		Sum2 = SumOfFibNumbers(100)
		wg.Done()
	}()

	sum3 := SumOfFibNumbers(20)
	println(sum3)

	wg.Wait()
	println(Sum1)
	println(Sum2)

}

func SumOfFibNumbers(r uint) int {
	sum := 0
	a, b := 0, 1
	for i := uint(1); i <= r; i++ {
		sum += a
		a, b = b, a+b
	}
	return sum
}
