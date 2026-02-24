package main

import (
	"fmt"
	"math/rand/v2"
)

func main() {

	slice1 := make([]int, 2, 4)
	slice1[0], slice1[1] = 10, 20

	slice2 := slice1 // it is only the header copy

	slice2[0] = 100

	fmt.Println(slice1)

	slice2 = append(slice2, 30)

	slice2[1] = 200

	PrintSliceHeader("slice1", slice1)
	PrintSliceHeader("slice2", slice2)
	fmt.Println(slice1)

	slice2[2] = 300
	fmt.Println(slice1) // no change in slice1

	fmt.Println(slice2)
	slice2 = append(slice2, 40, 50, 60, 70, 80) // len 8

	slice2 = append(slice2, 90) // cap 16

	PrintSliceHeader("slice2", slice2)
	slice2[0] = 10000
	fmt.Println(slice1)

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
