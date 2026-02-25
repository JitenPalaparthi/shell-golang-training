package main

import (
	"fmt"
	"unsafe"
)

func main() {

	slice1 := []int{}

	// 8 8 8 --> 24 bytes

	ptr := uintptr(unsafe.Pointer(&slice1))

	//fmt.Printf("%d 0x%x %p\n", ptr, ptr, &slice1)

	// t1 := T1{a: 100, b: 200, c: 300, d: 400.5}

	// fmt.Printf(" %p\n", &t1)

	// ptr1 := uintptr(unsafe.Pointer(&t1))

	// fmt.Printf("%d 0x%x\n", ptr1, ptr1)

	// ptr1 += 8

	// v := (*int)(unsafe.Pointer(ptr1))

	// println(*v)

	v := (*int)(unsafe.Pointer(ptr))
	//println(*v)
	fmt.Printf("0x%x\n", *v)

	ptr += 8

	v = (*int)(unsafe.Pointer(ptr))
	println(*v)

	ptr += 8
	v = (*int)(unsafe.Pointer(ptr))
	println(*v)

	var slice2 []int //:= []int{}

	ptr2 := uintptr(unsafe.Pointer(&slice2))

	v2 := (*int)(unsafe.Pointer(ptr2))
	//println(*v2)
	fmt.Printf("0x%x\n", *v2)

	var str1, str2 string

	ptr3, ptr4 := uintptr(unsafe.Pointer(&str1)), uintptr(unsafe.Pointer(&str2))

	v3, v4 := (*int)(unsafe.Pointer(ptr3)), (*int)(unsafe.Pointer(ptr4))

	fmt.Printf("0x%x\n", *v3)
	fmt.Printf("0x%x\n", *v4)

	arr := [2]int{10, 20}

	ptr5 := (*int)(unsafe.Pointer(&arr))

	println(*ptr5)

	arr1 := [0]int{}

	fmt.Printf("%p\n", &arr1)
	//1001acec0 B runtime.zerobase
	fmt.Printf("%p\n", &T2{})
}

type T2 struct{}
type T1 struct {
	a int
	b int
	c int
	d float64
}
