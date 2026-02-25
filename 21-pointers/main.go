package main

import (
	"fmt"
	"math/rand/v2"
)

func main() {
	var num1, ok1 = 100, true

	var ptr1, ptr2 = &num1, &ok1 // reference
	// var (
	// 	ptr1 *int
	// 	ptr2 *bool
	// )
	// both of them are 8 bytes on 64bit

	// var ptr1 *int
	// var ptr2 *bool

	*ptr1 = 200   // dereference
	*ptr2 = false // dereference

	println(*ptr1, *ptr2) // dereference

	var ptr3 *bool // there is a pointer, which is suppose to hold the address but the address is nil

	if ptr3 == nil {
		println("ptr3 is nil")
	}

	ptr4 := new(int) // what is this? no variable
	// new allocates memory ? where -> stack or heap

	println(*ptr4)

	*ptr4 = rand.IntN(1000)

	ptr5 := new([5]int)
	fmt.Println((*ptr5)[0])

	// can I create a poiner of a slice

	ptr6 := new([]int) // creates a pointer of 24 bytes

	fmt.Println(*ptr6)

	*ptr6 = append(*ptr6, 10, 11, 12, 13, 14, 15)
	fmt.Println(*ptr6)

	var arr_ptr *[5]int = &[5]int{10, 11, 12, 13, 14}
	var ptr_arr [2]*int

	a, b := 10, 20
	ptr_arr[0] = &a
	ptr_arr[1] = &b

	sum := 0
	for i := 0; i < len(*arr_ptr); i++ {
		sum += (*arr_ptr)[i]
	}
	println("sum of ptr_arr", sum)

	sum = 0

	for i := 0; i < len(ptr_arr); i++ {
		sum += *ptr_arr[i]
	}

	println("sum of arr_ptr", sum)

}

// ptr of string --> dummy pointer --? not is true 0x00 (nil)
// strings are considered as values --> ""

// pointer holds the address
//
