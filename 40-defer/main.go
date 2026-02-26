package main

func main() {
	str := "Hello World!"
	for _, c := range str {
		defer print(string(c))
	}
	a := 100
	defer println("defer func", a) // at a lapter point
	a = a + 1
	println("normal execution", a)
	p := new(int)
	*p = 100

	defer func(ptr *int) {
		println("defer pointer", *ptr)
	}(p)

	*p = *p + 1

	println("normal execution", *p)

	func() { //func2
		defer println("hello World -->defer func1")
	}()
}
