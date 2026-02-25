package main

type IShape interface {
	// IArea
	// IPerimeter
	Area() float64
	Perimeter() float64
	What() string
}

type IArea interface {
	Area() float64
}

type IPerimeter interface {
	Perimeter() float64
}
