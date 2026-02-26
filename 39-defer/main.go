package main

func main() {
	defer println("End of main")
	println("start of main")

	defer func() { //func1
		defer println("End of func1")
		println("start of func1")
		func() { //func1.1
			for i := 1; i <= 10; i++ {
				println(i)
			}
		}()
	}()

}

// defer defers the execution to the end of the call stack
