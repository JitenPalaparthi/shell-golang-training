package main

import (
	"errors"
	"fmt"
	"reflect"
)

func main() {
	var any1 any = uint8(100)
	// 	if v, ok := any1.(int); !ok {
	// 		if v1, ok1 := any1.(int8); !ok1 {
	// 			if v2, ok2 := any1.(uint8); !ok2 {
	// 				////
	// 			} else {
	// 				fmt.Println("uint8", v2)
	// 			}
	// 		} else {
	// 			fmt.Println("int8", v1)
	// 		}
	// 	} else {
	// 		fmt.Println("int", v)
	// 	}

	any1 = "Hello World"
	switch v1 := any1.(type) { // This gives you the type
	case uint:
		fmt.Println(v1)
	case uint8:
		fmt.Println(v1)
	case uint16:
		fmt.Println(v1)
	case uint32:
		fmt.Println(v1)
	case uint64:
		fmt.Println(any1.(uint64))
	case int:
		fmt.Println(any1.(int))
	case int8:
		fmt.Println(any1.(int8))
	case int16:
		fmt.Println(any1.(int16))
	case int32:
		fmt.Println(any1.(int32))
	case int64:
		fmt.Println(any1.(int64))
	case float32:
		fmt.Println(any1.(float32))
	case float64:
		fmt.Println(any1.(float64))
	default:
		fmt.Println("It is not a number")
	}

	fmt.Println("calling add function")

	if sum, err := Add(10, 20); err != nil {
		fmt.Println(err.Error())
	} else {
		fmt.Println("Sum:", sum)
	}

	if sum, err := Add(true, 20); err != nil {
		fmt.Println(err.Error())
	} else {
		fmt.Println("Sum:", sum)
	}

	if sum, err := Add("hello", "World"); err != nil {
		fmt.Println(err.Error())
	} else {
		fmt.Println("Sum:", sum)
	}

	if sum, err := Add(uint8(255), uint8(255)); err != nil {
		fmt.Println(err.Error())
	} else {
		fmt.Println("Sum:", sum)
	}

	if sum, err := Add(float32(255.34), float32(255.22413)); err != nil {
		fmt.Println(err.Error())
	} else {
		fmt.Println("Sum:", sum)
	}

	sum := AddG1(10, 20)
	fmt.Println("Sum:", sum)

	sum = AddG2(10, 20)
	fmt.Println("Sum:", sum)

	sum1 := AddG2(10.123, 20.123)
	fmt.Println("Sum:", sum1)

	// sum2 := AddG2(true, false)
	// fmt.Println("Sum:", sum1)

}

func IsNumber(a any) bool { // not asserting but checking the type using type switch
	switch a.(type) {
	case int, int8, int16, int32, int64, uint, uint8, uint16, uint32, uint64, float32, float64:
		return true
	default:
		return false
	}
}

func AreSame(a, b any) bool {
	if reflect.TypeOf(a) == reflect.TypeOf(b) {
		return true
	}
	return false
}

func Add(a, b any) (any, error) {
	if !AreSame(a, b) {
		return 0, fmt.Errorf("first and second args are different")
	}

	if !IsNumber(a) {
		return 0, errors.New("first arg is not a number")
	}
	if !IsNumber(b) {
		return 0, errors.New("second arg is not a number")
	}

	switch v1 := a.(type) { // This gives you the type
	case uint:
		return v1 + b.(uint), nil
	case uint8:
		return uint16(v1 + b.(uint8)), nil
	case uint16:
		return uint32(v1 + b.(uint16)), nil
	case uint32:
		return uint64(v1 + b.(uint32)), nil
	case uint64:
		return v1 + b.(uint64), nil
	case int:
		return v1 + b.(int), nil
	case int8:
		return int16(v1 + b.(int8)), nil
	case int16:
		return v1 + b.(int16), nil
	case int32:
		return int64(v1 + b.(int32)), nil
	case int64:
		return v1 + b.(int64), nil
	case float32:
		return float64(v1 + b.(float32)), nil
	case float64:
		return v1 + b.(float64), nil
	default:
		return 0, errors.New("some error")
	}
}

// >= 1.18

func AddG1[T int | int8 | int16 | int32 | int64 | uint | uint8 | uint16 | uint32 | uint64 | float32 | float64](a, b T) T {
	return a + b
}

func AddG2[T NumberType](a, b T) T {
	return a + b
}

type NumberType interface {
	int | int8 | int16 | int32 | int64 | uint | uint8 | uint16 | uint32 | uint64 | float32 | float64
}
