package main

import (
	"fmt"
	"math/rand/v2"
	"reflect"
)

func main() {

	//var val int // 0
	//const PI float32 = 3.14

	var arr1 [5]int // there is a zero value for the array
	fmt.Println("values:", arr1, "Type;", reflect.TypeOf(arr1))

	for i := 0; i < len(arr1); i++ {
		arr1[i] = rand.IntN(100) // mutate
	}
	fmt.Println("values:", arr1, "Type;", reflect.TypeOf(arr1))

	arr2 := [5]int{10, 11, 12, 13, 14} // assigning array with values short hand declaration
	arr3 := [...]int{10, 11, 12}       // length is 3 by the compiler

	sum := 0
	for _, v := range arr2 {
		sum += v
	}
	println("sum:", sum)

	sum = SumOf(arr2)
	println("sum:", sum)

	// sum = SumOf1(arr3)
	// println("sum:", sum)

	sum = SumofAny(arr1)
	println("sum:", sum)
	sum = SumofAny(arr2)
	println("sum:", sum)
	sum = SumofAny(arr3)
	println("sum:", sum)

	arr4 := [...]int{10, 11, 12, 13}
	sum = SumofAny(arr4)
	println("sum:", sum) //it gives 0 bcz not implemented for [4]int

	//arr5 := [5]int(arr4) // cannot type cast
}

// arrays are fixed length
// arrays can be on stack or heap , it depends on the compiler's escape analysis
// array does not maintain any kind of header, so it understand the length purely based on the type
// cannot extend array
// dont use arrays as input params for functions, due to operational issues..
// cannot type cast between arrays of different types

func SumOf(arr [5]int) int {
	sum := 0
	for _, v := range arr {
		sum += v
	}
	return sum
}

// generally you dont use array as an input param to the function
func SumOf1(arr [3]int) int {
	sum := 0
	for _, v := range arr {
		sum += v
	}
	return sum
}

func SumofAny(a any) int {
	sum := 0
	switch arr := a.(type) {
	case [3]int:
		for _, v := range arr {
			sum += v
		}
		return sum
	case [5]int:
		sum = SumOf(arr)
		return sum
	case [10]int:
		for _, v := range arr {
			sum += v
		}
		return sum
	}
	return sum
}
