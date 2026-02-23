package main

import (
	"fmt"
	"reflect"
	"strconv"
)

// package main is create and bootstraped by go runtime
// no need to install any runtime, it is embedded with every binary of Go

// there is not implicit in go
func main() {

	var num1 uint8 = 100

	var num2 uint64 = uint64(num1)

	// (uint64)num1

	println(num2)

	// all numbers can be casted among .. but bool cannot be to num

	var num3 uint64 = 989898989
	fmt.Printf("%b\n", num3) // 1110110000000010101000 11101101
	//												   11101101

	var num4 uint8 = uint8(num3)

	fmt.Printf("%d-> %b\n", num4, num4)

	float1 := float32(123.23)

	var num5 uint32 = uint32(float1)
	fmt.Println(num5)

	num6 := 1231231231231
	float2 := float64(num6)

	fmt.Printf("%.2f\n", float2)

	a, b := 10, 20

	var c int
	c = int(add(uint32(a), uint32(b)))

	println(c)

	ok1 := true

	//var b1 uint8 = uint8(ok1)

	var b1 uint8 = BoolToUint8(ok1)

	println(b1)

	// rune

	var char1 rune = 'A' // compiler converst the 'A' to its utf8 (4 bytes) value

	char1 = char1 + 1
	char1 = char1 + 100
	char1 = char1 - 100

	var char2 uint8 = 'B'
	var char3 int32 = '😀'

	var char4 uint64 = '😀'

	fmt.Println(char1, char2, char3, char4)

	//gender := rune('F')

	// if gender == 'F' {

	// }

	//var i integer = 12312

	println(string(char3)) // type casting of a char

	println(string(128512))

	var num7 uint64 = uint64(char1)

	println(string(num7))

	//var num8 int32 = 66

	// if uint64(char1) == uint64(num8) {

	// }

	a1, s1, m1, sb1 := calc(10, 20)

	println(a1, s1, m1, sb1)

	a2, _, _, _ := calc(10, 20) // _ blank identifier

	println(a2)
	_, _, _, su2 := calc(20, 10)
	println(su2)

	_, d := 100, 200 // _ blank identifier does not allocate any memory

	println(d)

	// string conversions

	var num8 int = 98983232

	str := strconv.Itoa(num8)

	//str = fmt.Sprint(num8)

	println(str)

	a3, b3, c3, d3 := "Hello Wrold", 123213, 123213.213, true

	str = fmt.Sprint("string:", a3, " int:", b3, " float64", c3, " bool:", d3)
	println("format:", str)

	str = "123123123123"

	i, _ := strconv.Atoi(str)
	fmt.Println(i, "-->", reflect.TypeOf(i))

	str = "123123ede123123" // this cannot be converted

	i, err := strconv.Atoi(str)
	if err != nil { //  if not nil means there is some error
		println(err.Error()) // parsing "1231a23123123": invalid syntax
	} else {
		fmt.Println(i, "-->", reflect.TypeOf(i))
	}
	//true --> 1, t, T, TRUE, true, True,
	// false --> 0, f, F, FALSE, false, False.

	ok2, err := strconv.ParseBool("u")
	if err != nil {
		println(err.Error())
	} else {
		println(ok2)
	}

	str = "no"
	ok2, err = ParseBool(str)

	if err != nil {
		println(err.Error())
	} else {
		println(ok2)
	}

}

// strconv.ParseFloat
// strconv.ParseInt

//type integer = int

func BoolToUint8(b bool) uint8 {
	if b {
		return 1
	}
	return 0
}

func ParseBool(s string) (bool, error) {
	ok, err := strconv.ParseBool(s)
	if err == nil {
		return ok, nil
	} else {
		if s == "YES" || s == "Yes" || s == "yes" || s == "Y" || s == "y" {
			return true, nil
		} else if s == "NO" || s == "No" || s == "no" || s == "N" || s == "n" {
			return false, nil
		} else {
			return false, fmt.Errorf("parsing %v invalid syntax", s)
		}
	}
}

func add(a, b uint32) uint32 {
	return (a + b)
}

func calc(a int, b int) (add int, sub int, mul int, div int) {
	//func calc(a, b int) (int, int,int,int) {
	add, sub = a+b, a-b
	return add, sub, a * b, a / b
	//return a + b, a - b, a * b, a / b
}

// type cast
// type conversion
// type assertion
