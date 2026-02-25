package main

func main() {

	sum := calc(10, 20, func(i1, i2 int) int {
		return i1 + i2
	})

	println("sum:", sum)

	sb := calc(20, 10, sub)

	println("sub:", sb)

	mul := func(a, b int) int {
		return a * b
	}

	ml := calc(10, 20, mul)
	println("mul:", ml)

	m1 := make(map[string]func(int, int) int)

	a, b := 20, 10

	m1["add"] = func(i1, i2 int) int {
		return i1 + i2
	}

	m1["sub"] = sub

	m1["mul"] = mul

	for k, fn := range m1 {
		println("executing a func>>", k)
		r := fn(a, b)
		println("result of ", k, "=", r)
	}

	fn := eval(20, 10, "sub")
	if fn != nil {
		r := fn()
		println("r:", r)
	} else {
		println("fn is nil")
	}

	fn = eval(20, 10, "add")
	if fn != nil {
		r := fn()
		println("r:", r)
	} else {
		println("fn is nil")
	}

	fn = eval(20, 10, "mul")
	if fn != nil {
		r := fn()
		println("r:", r)
	} else {
		println("fn is nil")
	}

	fn = eval(20, 10, "something")
	if fn != nil {
		r := fn()
		println("r:", r)
	} else {
		println("fn is nil")
	}

}

func calc(a, b int, fn func(int, int) int) int {
	return fn(a, b)
}

func sub(a, b int) int {
	return a - b
}

func eval(a, b int, fn string) func() int {
	switch fn {
	case "add", "sum":
		return func() int {
			return a + b
		}
	case "sub":
		return func() int {
			return a - b
		}
	case "mul":
		return func() int {
			return a * b
		}
	default:
		return nil
	}
}
