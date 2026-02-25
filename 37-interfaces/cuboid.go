package main

type Cuboid struct {
	L, B, H float32
}

func NewCuboid(l, b, h float32) *Cuboid {
	return &Cuboid{l, b, h}
}

func (r *Cuboid) Area() float64 {
	return float64(r.L * r.B * r.H)
}

func (r *Cuboid) Perimeter() float64 {
	return 2 * float64(r.L+r.B+r.H)
}

func (*Cuboid) What() string { // can avvoid receiver name
	return "Cuboid"
}
