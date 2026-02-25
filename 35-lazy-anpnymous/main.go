package main

func main() {

	slice1 := []func(){}

	// 1.22< the loop variable is a global variable for the loop
	// 1.22=> the loop variable is a new variable for the each iteration
	for i := 1; i <= 10; i++ {
		slice1 = append(slice1, func() {
			println(&i, i)
		})
	}

	for _, fn := range slice1 {
		fn()
	}

	i := 1 // global
	slice2 := []func(){}
loop:
	i++
	slice2 = append(slice2, func() {
		println(&i, i)
	})
	if i <= 10 {
		goto loop
	}

	for _, fn := range slice2 {
		fn()
	}

	// global with local call

	i = 1 // global
	slice3 := []func(int){}
loop1:
	i++

	slice3 = append(slice3, func(a int) {
		println(a)
	})

	if i <= 10 {
		goto loop1
	}

	for i, fn := range slice3 {
		fn(i + 1)
	}

}
