package main

func main() {

	str1 := "Hello World"

	for i := 0; i < len(str1); i++ {
		print(string(str1[i]), "--")
	}
	println()

	str2 := "Hello 你好"

	for i := 0; i < len(str2); i++ {
		print(string(str2[i]), "--")
	}
	println()

	for i, v := range str2 { // both index and value
		println("index:", i, "value:", string(v))
	}

	for _, v := range str2 { // only value
		println("value:", string(v))
	}

	//for i, _ := range str2 { // only index
	for i := range str2 { // only index
		println("index", i)
	}

	for i := range 10 { // 0-9, works only go v1.22 onwards
		if i%2 == 0 {
			print(i, " ")
		}
	}

}

// range loop
// range loop is different for differnet constructs
// on string, array,slice --> index and value
// on maps --> key, value
// on channel --> receive value from the channel
