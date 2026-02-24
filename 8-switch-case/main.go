package main

import "math/rand/v2"

func main() {

	day := uint8(4)

	switch day {
	case 1:
		println("Sunday")
	case 2:
		println("Monday")
	case 3:
		println("Tuersday")
	case 4:
		println("Wednesday")
	case 5:
		println("Thursday")
	case 6:
		println("Friday")
	case 7:
		println("Saturday")
	default:
		println("no day")
	}

	switch char := GetChar(); char {
	case 'a', 'e', 'i', 'o', 'u', 'A', 'E', 'I', 'O', 'U':
		println(string(char), "is vovel")
	default:
		{
			println(string(char), "is cosonent or special char")
		}
	}

	// empty switch

	num := GetNum()

	switch { // empty switch, for empty switch can write conditons in cases
	case num >= 0 && num <= 50:
		println(num, "is between 0-50")
	case num > 50 && num <= 100:
		println(num, "is between 51-100")
	default:
		println(num, "is > 100")
	}

	println("using fallthrough")

	num = 40
	// to achieve samething in other programming removing the break, you need use
	// fallthrough

	switch {
	case num%8 == 0:
		println(num, "is divisible by 8")
		fallthrough
	case num%4 == 0:
		println(num, "is divisible by 4")
		fallthrough
	case num%2 == 0:
		println(num, "is divisible by 2")
	}

	println("false negative")

	num = 6
	switch {
	case num%2 == 0:
		println(num, "is divisible by 2")
		fallthrough
	case num%4 == 0:
		println(num, "is divisible by 4")
		fallthrough
	case num%8 == 0:
		println(num, "is divisible by 8")
	}

}

func GetChar() byte {
	return uint8(rand.IntN(256))
}

func GetNum() byte {
	return uint8(rand.IntN(256))
}
