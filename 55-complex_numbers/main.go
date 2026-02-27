package main

import "fmt"

func main() {
	c1 := 10 + 2.4i // real and imag

	r := real(c1)
	i := imag(c1)

	println(r, i)

	c2 := complex(12.3, 2) // complex128

	r1, i1 := float32(1.1), float32(0.2)
	c3 := complex(r1, i1) // complex64

	c4, c5 := c2+complex128(c3), c2-complex128(c3)
	fmt.Println(c4, c5)
}
