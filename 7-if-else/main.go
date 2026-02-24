package main

import (
	"math/rand/v2"
	"strconv"
)

// package main is create and bootstraped by go runtime
// no need to install any runtime, it is embedded with every binary of Go

func main() {
	//a, b := 10, 20

	//c := (a + b) / (b - a) * 10 * (a) / 10 // operator precedence

	str := "12312234e"

	// v, err := strconv.Atoi(str)

	// if err != nil {
	// 	println(err.Error())
	// } else {
	// 	println(v)
	// }

	// var v int
	// var err error
	if v, err := strconv.Atoi(str); err != nil {
		println(err.Error())
	} else {
		println(v)
	}

	// println(v)

	if age, gender := uint8(41), uint8('m'); age >= 21 && gender == 'M' || gender == 109 {
		println("he is eligible for marriage")
	} else if age >= 18 && gender == 'F' || gender == 'f' {
		println("she is eligible for marriage")
	} else {
		println("not eligible")
	}

	println("get magic age and gender")
	if age, gender := GetAgeNGender(); age >= 21 && gender == 'M' || gender == 109 {
		println("he is eligible for marriage because of age", age, " gender:", string(gender))
	} else if age >= 18 && gender == 'F' || gender == 'f' {
		println("she is eligible for marriage because of age", age, " gender:", string(gender))
	} else {
		println("not eligible for marriage because of age", age, " gender:", string(gender))
	}

	//var char uint8 = '😀'

}

func GetAgeNGender() (uint8, byte) {
	age := uint8(rand.IntN(50))
	if age%2 == 0 {
		return age, uint8('M')
	} else {
		return age, byte('F')
	}
}

// arth
// comp
// logical
// bit
