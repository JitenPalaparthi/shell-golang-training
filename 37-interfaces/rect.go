package main

type Rect struct {
	L, B float32
}

func NewRect(l, b float32) *Rect {
	return &Rect{l, b}
}

func (r *Rect) Area() float64 {
	return float64(r.L * r.B)
}

func (r *Rect) Perimeter() float64 {
	return 2 * float64(r.L+r.B)
}

func (*Rect) What() string { // can avvoid receiver name
	return "Rect"
}
