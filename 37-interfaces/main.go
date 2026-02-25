package main

import (
	"fmt"
	"math/rand/v2"
)

func main() {

	slice := []IShape{} // a slice can store any variable type which has implemted IShape interface

	slice = append(slice, NewRect(10.2, 12.3), NewSquare(13.4), NewRect(23.45, 12.34), NewSquare(87.5))

	slice = append(slice, NewCuboid(12, 13, 14), NewCuboid(12.3, 14.5, 21.9))

	for _, s := range slice {
		Shape(s)
		println()
	}

	for i := 0; i < 5; i++ {
		r := rand.IntN(len(slice))
		Shape(slice[r])
		println()
	}

}

func Shape(ishape IShape) {
	fmt.Println("Area of ", ishape.What(), ":", ishape.Area())
	fmt.Println("Perimeter of ", ishape.What(), ":", ishape.Perimeter())
}

// a table itable
// for each concrete type (Rect,Square,Cuboid)
// A new Itable is created and all the repective method pointers are stored there

// 15 lines
// Runtime Poly by virtual tables/ itables
