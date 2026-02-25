package main

import (
	"errors"
	"fmt"
)

func main() {

	slice1 := []int{10, 11, 12}

	sum, slice1 := SumOfAddElemR(slice1, 13, 14, 15, 16, 17) // call by value of slice

	println(sum)

	if sum, err := SumOfAddElemPtr(&slice1, 100, 200, 300, 123, 123, 34, 98, 56, 76); err != nil {
		println(err.Error())
	} else {
		println(sum)
	}

	var slice2 []int // nil

	if slice2 == nil {

	}

	if sum, err := SumOfAddElemPtr(&slice2, 100, 200, 300, 123, 123, 34, 98, 56, 76); err != nil {
		println(err.Error())
	} else {
		println(sum)
	}

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

func SumOfAddElemPtr(slice *[]int, nums ...int) (int, error) {
	if *slice == nil { // since nil can be checked on normal slice itself, to check the pointer inside the slice we need to use *slice==nil
		return 0, errors.New("nil slice pointer")
	}
	sum := 0
	*slice = append(*slice, nums...) // when ever you use appent inside a func, that would not replicate outside slice
	for _, v := range *slice {
		sum += v
	}
	fmt.Println("insidde func:", slice)
	return sum, nil
}
