package main

import (
	"fmt"
	"math/rand/v2"
)

func main() {

	var slice1 []int // only the declaration of slice but not instantiated

	if slice1 == nil {
		println("nil slice")
		//slice1 = make([]int, 5) // without cap so len= cap
		slice1 = make([]int, 5, 10) // with cap so len=5 cap=10
	}

	fmt.Println(slice1)

	slice1[0], slice1[1], slice1[2], slice1[3] = 10, 11, 12, 13
	fmt.Println(slice1)

	// sum := 0

	// for _, v := range slice1 {
	// 	sum += v
	// }

	sum := SumOf(slice1)

	println("sum:", sum)

	slice2 := make([]int, 10) // len 10 cap 10

	FillSlice(slice2)

	fmt.Println(slice2)

	sum = SumOf(slice2)

	println("sum:", sum)

	slice3 := []int{10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20, 21, 22}

	sum = SumOf(slice3)
	println("sum:", sum)

	slice5 := []int{} // it is not a nil slice
	// ptr if the slice is some dummy ptr
	if slice5 == nil {
		println("nil slice")
	} else {
		fmt.Println(slice5)
	}

	arr1 := [23]int{10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20, 21, 22}

	sum = SumOf(arr1[:]) // array can be converted to the slice
	println("sum:", sum)

	slice4 := arr1[:] // array to slice ptr:arr1[0],len:23,cap:23

	sum = SumOf(slice4)
	println("sum:", sum)
	slice4[0] = 99999 // when change slice4, it chances arr1 as well
	fmt.Println(arr1)

	slice6 := []int{1, 2, 3, 4, 5}

	slice7 := slice6 // when a slice to another slice , the header of the slice6 is copied to slice7

	slice7[0] = 9999
	fmt.Println(slice6)

	PrintSliceHeader("slice6", slice6)
	PrintSliceHeader("slice7", slice7)

	slice7 = append(slice7, 6) // extending the slice

	slice7[1] = 8888
	fmt.Println(slice6)
	fmt.Println(slice7)

	PrintSliceHeader("slice6", slice6)
	PrintSliceHeader("slice7", slice7)

}

// slice can be extended
// slice has ptr, len and cap
// slice size if 24 bytes
// slice data is stored in stack or heap

// make is a built in func--> make allocates memory? stack or heap based on ..
// zero value for the slice

// generally slice is used as an parm for a function , very rarely array is used

func SumOf(slice []int) int {
	sum := 0
	for _, v := range slice {
		sum += v
	}
	return sum
}

func FillSlice(slice []int) {
	for i := range slice {
		slice[i] = rand.IntN(1000)
	}
}

func PrintSliceHeader(name string, slice []int) {
	println(name, "details")
	println("---------------------")
	fmt.Printf("address of the slice:%p\n", &slice)
	fmt.Println("Slice:", slice)

	if len(slice) > 0 {
		fmt.Printf("Ptr :%p\n", &slice[0])
		fmt.Printf("Len :%d\n", len(slice))
		fmt.Printf("Cap :%d\n", cap(slice))

	}
	println("--------------------------\n")
}
