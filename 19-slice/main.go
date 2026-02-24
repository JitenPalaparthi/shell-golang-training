package main

import "fmt"

func main() {
	slice1 := []int{10, 11, 12, 13, 14, 15, 16, 17, 18, 19}

	slice2 := slice1

	slice3 := slice1[:5]

	slice4 := slice1[3:8]

	slice5 := slice1[5:]

	PrintSliceHeader("slice1", slice1)
	PrintSliceHeader("slice2", slice2)
	PrintSliceHeader("slice3", slice3)
	PrintSliceHeader("slice4", slice4)
	PrintSliceHeader("slice5", slice5)

	// copy is a built in function, which is deep copy of the slice, not a header copy

	slice6 := make([]int, 4)

	copy(slice6, slice1) // len slice6 is 4 , len slice1 =10

	PrintSliceHeader("slice6", slice6)

	slice7 := make([]int, 20)
	copy(slice7, slice1) // slice7 = 20 , slice1 =10
	// 10 elements gets copied and the rest 10 elements would be zeros
	PrintSliceHeader("slice7", slice7)

	slice8 := make([]int, 5)
	CopySlice(slice8, slice1)
	PrintSliceHeader("slice8", slice8)

	slice9 := make([]int, 0)
	if slice9 == nil {
		println("nil")
	} else {
		fmt.Println(slice9)
	}

	var slice10 []int // slice is nil

	slice10 = append(slice10, 10, 20, 30) // it allocates memoruy even it is nil and append values
	fmt.Println(slice10)

	slice11 := append(slice1[:4], slice1[5:]...)
	fmt.Println(slice11)

	arr := [5]int{100, 200, 300, 400, 500}

	slice12 := append(slice1, arr[:]...)
	fmt.Println(slice12)
	clear(slice12)
	fmt.Println(slice12)
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

func CopySlice(dest, src []int) {
	for i := 0; i < min(len(dest), len(src)); i++ {
		dest[i] = src[i]
	}
	//return dest
}

// write a func to delete a specifi element from the slice and return it
