package main

import (
	"fmt"
)

func main() {

	var (
		num  int    = 100
		num1 MyInt1 = 200
		num2 MyInt2 = 300
		num3 MyInt3 = 400
		//ok1  bool   = true
	)

	// num call string
	// num call sq
	// num call cube

	str1 := MyInt1(num).ToString()
	s1 := MyInt2(num).Sq()
	c1 := MyInt3(num).Cube()
	fmt.Println("Str:", str1, "Sq:", s1, "Cube:", c1)

	str2 := num1.ToString()
	s2 := MyInt2(num1).Sq()
	c2 := MyInt3(num1).Cube()
	fmt.Println("Str:", str2, "Sq:", s2, "Cube:", c2)

	str3 := MyInt1(num3).ToString()
	s3 := MyInt2(num3).Sq()
	c3 := num3.Cube()
	fmt.Println("Str:", str3, "Sq:", s3, "Cube:", c3)

	var float1 float32 = 12321.4343

	str4 := MyInt1(float1).ToString()
	s4 := MyInt2(float1).Sq()
	c4 := MyInt3(float1).Cube()
	fmt.Println("Str:", str4, "Sq:", s4, "Cube:", c4)

	str5 := MyInt1(num2).ToString()
	s5 := num2.Sq()
	c5 := MyInt3(num2).Cube()
	fmt.Println("Str:", str5, "Sq:", s5, "Cube:", c5)

	num1.Incr()

	println(num1)

	var myint1 MyInt1 = MyInt1(num)
	var ptr *MyInt1 = &myint1

	//var ptr1 *MyInt1 = MyInt1(new(int))

	ptr.Incr()
	fmt.Println(myint1)

	(*MyInt1)(&num).Incr()
	fmt.Println(num)

	var ptr1 = new(int)
	*ptr1 = 99 // is. apointer

	s7 := MyInt2(*ptr1).Sq()
	println(s7)
	(*MyInt1)(ptr1).Incr() // a poiinter of int can dictly be cased to pointer of MyInt1

}

type MyInt1 int // new type

type MyInt2 int //

type MyInt3 int

type Integer = int // alias , not a new type

func (m MyInt1) ToString() string {
	return fmt.Sprint(m)
}

func (m MyInt2) Sq() int {
	return int(m * m)
}

func (m MyInt3) Cube() int {
	return int(m * m * m)
}

func (m *MyInt1) Incr() {
	if m != nil {
		*m++
	}
}
