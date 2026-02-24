package main

func main() {

	//var nums ...int // cannot do this

	sum := SumOf()
	println("Sum:", sum)

	sum = SumOf(10, 20)
	println("Sum:", sum)

	sum = SumOf(10, 230, 23, 35, 65, 7, 8, 9, 5, 6, 67, 4, 3, 46, 57, 566, 3, 34, 5)
	println("Sum:", sum)

	sumofMul := SumOfMul(2)
	println("Sum:", sumofMul)

	sumofMul = SumOfMul(2, 10, 20)
	println("Sum:", sumofMul)

	slice1 := []int{10, 230, 23, 35, 65, 7, 8, 9, 5, 6, 67, 4, 3, 46, 57, 566, 3, 34, 5}

	sum = SumOf(slice1...) // pass a slice as a variadic arg
	println("Sum:", sum)
	arr1 := [10]int{10, 20, 30, 40, 50, 60, 70, 80, 90}

	sum = SumOf(arr1[:]...)
	println("Sum:", sum)
}

func SumOf(nums ...int) int {
	sum := 0
	for _, v := range nums {
		sum += v
	}
	return sum
}

// func SumOfMul(nums ...int,mul int) int { // this does not work, becase the variadic must be the last one
func SumOfMul(mul int, nums ...int) int { // the variadic must be the last one
	sum := 0
	for _, v := range nums {
		sum += v * mul
	}
	return sum
}

// Variadic parameter must be the last parameter
// Variadic parameter can only be used in funcs and methods
// For variadic parameter, the arguments can be zero or any number of(based on memory)
// can use range loop on variadic
// can pass slice as a argument to the variaoc
// can pass arr as an argument to the variadic paramater but after converting the array to then slice
