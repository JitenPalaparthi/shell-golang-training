package main

import (
	"fmt"
	"unsafe"
)

func main() {
	arr1 := [5]int{10, 20, 30, 40, 50}
	ptr := &arr1 // address of array is nothing but the 0th element of an array
	ptr1 := &arr1[0]
	fmt.Printf("%d %p\n", *ptr, ptr)   // developer friendly println
	fmt.Printf("%d %p\n", *ptr1, ptr1) // developer friendly println

	// ptr1 += 8 // pointer arthimetic is very dangerous operation

	// ptr
	// len

	// uintptr
	// unsafe.Pointer

	// cannot do poiner arth on pointers, but can do on numbers

	ptr2 := uintptr(unsafe.Pointer(&arr1)) // 1 and 4 here
	println(ptr2)
	ptr2 += 8

	v := (*int)(unsafe.Pointer(ptr2)) // 3 and 2

	println(*v)

	println("array with pointers")
	for i := 0; i < len(arr1); i++ {
		ptr2 := uintptr(unsafe.Pointer(&arr1[i])) //
		v := (*int)(unsafe.Pointer(ptr2))
		println(*v)
		ptr2 += unsafe.Sizeof(arr1[0])
	}

	str := "Hello World"

	println(*(&str))

	ptr3 := uintptr(unsafe.Pointer(&str))
	ptr3 += 8

	ptr_len := (*int)(unsafe.Pointer(ptr3))
	*ptr_len = 100

	println(len(str))

	// ptr4 := new(int)

	// println(ptr)
	// println(&ptr)
	println(str)
}

// 0x2c9b72738000

// 1.A pointer value of any type can be converted to a Pointer.
// 2.A Pointer can be converted to a pointer value of any type.
// 3.A uintptr can be converted to a Pointer.
// 4.A Pointer can be converted to a uintptr.
