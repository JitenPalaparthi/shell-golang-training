package main

import "fmt"

func main() {
	r1 := New(10.1, 12.3)

	a1 := Area1(r1)
	println("Area of r1:", a1)

	fmt.Println(r1.A)

	a2 := Area2(&r1)
	println("Area of r1:", a2)

	fmt.Println(r1.A)

	r2 := New(10.1, 12.3)
	a3 := (&r2).Area()
	println("Area of r2:", a3)
	fmt.Println(r2.A)

	p3 := r2.Perimeter()
	println("Perimeter of r2:", p3)
	fmt.Println(r2.P)

	r3 := NewP(10.1, 12.3)
	a4 := r3.Area()
	println("Area of r3:", a4)
	fmt.Println(r3.A)

	p4 := r3.Perimeter()
	println("Perimeter of r3:", p4)
	fmt.Println(r3.P)

}

type Rect struct {
	L, B float32
	A, P float64
}

func New(l, b float32) Rect {
	return Rect{l, b, 0, 0} // Rect{L:l,B:b,A:0,P:0.0}
}

func NewP(l, b float32) *Rect {
	return &Rect{l, b, 0, 0} // Rect{L:l,B:b,A:0,P:0.0}
}

func Area1(r Rect) float64 { // call by value
	r.A = float64(r.L * r.B)
	return r.A
}

func Area2(r *Rect) float64 { // pass by ref
	r.A = float64(r.L * r.B)
	return r.A
}

// func (r Rect) Area() float64 { // call by value
// 	r.A = float64(r.L * r.B)
// 	return r.A
// }

func (r *Rect) Area() float64 { // call by ref
	r.A = float64(r.L * r.B)
	return r.A
}

func (this *Rect) Perimeter() float64 { // call by ref
	this.P = 2 * float64(this.L+this.B)
	return this.P
}

// 99% receivers are pointers
//
