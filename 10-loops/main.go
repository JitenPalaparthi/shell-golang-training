package main

import "math/rand/v2"

func main() {

	a, b := 0, 1

	for i := 1; i <= 10; i++ {
		print(a, " ")
		a, b = b, a+b
	}

	// for like while

	i := 1

	println("\neven numbers")
	for i <= 10 { // for loop but looks like while loop
		if i%2 == 0 {
			print(i, " ")
		}
		i++
	}

	i = 1

	println("\nodd numbers")
	for ; i <= 10; i++ {

		if i%2 == 0 {
			continue
		}
		print(i, " ")
	}
	println()
	for i, j := 1, 10; i <= j; i, j = i+1, j-1 {
		println(i, j)
	}

	println("nested loop inner break")

	for i := 1; i <= 5; i++ {
		for j := 1; j <= 5; j++ {
			if j == 3 {
				break
			}
			println("i", i, "j:", j)
		}
	}

	println("nested loop inner and outer break")

	done := false
	for i := 1; i <= 5; i++ {
		if done {
			break
		}
		for j := 1; j <= 5; j++ {
			if j == 3 {
				done = true
				break
			}
			println("i", i, "j:", j)
		}
	}

	println("nested loop inner and outer break using lable")

exit:
	for i := 1; i <= 5; i++ {
		for j := 1; j <= 5; j++ {
			for k := 1; k <= 5; k++ {
				if k == 3 {
					break exit
				}
				println("i:", i, "j:", j, "k:", k)
			}
		}
	}

	// what are the areas break is applicable
	// for , switch

	// unconditional loop

outer:
	for {
		num := rand.IntN(100)
		switch {
		case num%13 == 0:
			println(num, "is div by 13")
			break outer
		case num%2 == 0:
			{
				println(num, "even number")
			}
		default:
			println(num, "odd number")
		}
	}
}
