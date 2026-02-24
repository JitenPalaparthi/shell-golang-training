package main

import (
	"fmt"
	"math/rand/v2"
	"reflect"
)

func main() {
	arr2d := [2][3]int{{1, 2, 3}, {4, 5, 6}}
	arr3d := [2][2][3]int{{{1, 2, 3}, {4, 5, 6}}, {{1, 2, 3}, {4, 5, 6}}}

	fmt.Println("type of arr2d:", reflect.TypeOf(arr2d))
	fmt.Printf("type of arr3d:%T\n", arr3d)

	for _, arr1 := range arr2d {
		for _, v := range arr1 {
			print(v, " ")
		}
		println()
	}

	println("using normal loop")

	for i := 0; i < len(arr2d); i++ {
		for j := 0; j < len(arr2d[i]); j++ {
			print(arr2d[i][j], " ")
		}
		println()
	}

	println("getting values of 3d array")

	for _, arr1 := range arr3d {
		for _, arr2 := range arr1 {
			for _, v := range arr2 {
				print(v, " ")
			}
		}
		println()
	}

	var arr2d1 [3][3]int

	for i, arr1 := range arr2d1 {
		for j := range arr1 {
			arr2d1[i][j] = rand.IntN(100)
		}
	}

	fmt.Println(arr2d1)

}
