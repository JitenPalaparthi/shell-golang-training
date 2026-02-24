package main

import "fmt"

func main() {

	//var slice1 []int

	//slice2 := []int{} // not nil

	slice3 := []int{10, 20, 30} //len:3 cap:3

	sum := SumOfAddElem(slice3, 40, 50, 60)

	println("Sum:", sum)
	fmt.Println(slice3)

	sum, slice3 = SumOfAddElemR(slice3, 40, 50, 60)
	fmt.Println(slice3)
}

func SumOfAddElem(slice []int, nums ...int) int {
	sum := 0
	slice = append(slice, nums...) // when ever you use appent inside a func, that would not replicate outside slice
	for _, v := range slice {
		sum += v
	}
	fmt.Println("insidde func:", slice)
	return sum
}

func SumOfAddElemR(slice []int, nums ...int) (int, []int) {
	sum := 0
	slice = append(slice, nums...) // when ever you use appent inside a func, that would not replicate outside slice
	for _, v := range slice {
		sum += v
	}
	fmt.Println("insidde func:", slice)
	return sum, slice
}
