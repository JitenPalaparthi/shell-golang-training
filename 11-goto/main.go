package main

func main() {

	i := 1
loop:
	//println(i
	if i > 10 {
		goto exit
	}
	if i%2 == 0 {
		goto even
	} else {
		goto odd
	}

even:
	println(i, "is even")
	i++
	goto loop
odd:
	println(i, "is ood")
	i++
	goto loop

exit:
	println("exiting")
}
