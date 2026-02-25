package main

// to do the escape analysys
// go build -gcflags="-m" main.go

var GPtr *int = new(int) // global variable // generally in ds

func main() {

	a := 100

	incr(a) // call by value
	println(a)
	incrP(&a) // call by ref
	println(a)

	r := sq1(100)

	println(r)

	ptr1 := new(int) // does not escape means does not store in heap memory

	sq2(100, ptr1) // out

	println(*ptr1)

	ptr2 := sq3(100) // initially does not escape

	println(*ptr2)

	ptr3 := sq4(100)
	println(*ptr3)

	arr1 := [500]int{}
	arr2 := [100000]int{} // moved to heap ? why?
	println(len(arr1))
	println(len(arr2))

	slice1 := make([]int, 1000) // may be based on program , this also can be moved to heap
	slice2 := make([]int, 10000000)

	println(len(slice1))
	println(len(slice2))
}

func incr(n int) { // stack frame --> int
	n++
}

func incrP(p *int) {
	if p != nil {
		*p++ // deref and incrementing the value
	}
}

func sq1(n int) int {
	return n * n
}

func sq2(n int, out *int) {
	if out != nil {
		*out = n * n
	}
}

func sq3(n int) *int { // dangling pointer in c and c++ but go handles it gracefulluy by string p in heap
	p := new(int) // escapes to heap
	*p = n * n
	return p
}

func sq4(n int) *int {
	if GPtr != nil {
		*GPtr = n * n
		return GPtr
	}
	return nil
}
