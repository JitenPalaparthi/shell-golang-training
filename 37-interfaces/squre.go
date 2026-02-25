package main

type Square float32

func NewSquare(s float32) Square {
	return Square(s)
}

func (s Square) Area() float64 {
	return float64(s * s)
}

func (s Square) Perimeter() float64 {
	return 4 * float64(s)
}

func (Square) What() string {
	return "Square"
}
