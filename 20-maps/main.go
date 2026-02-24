package main

import (
	"errors"
	"fmt"
	"math/rand/v2"
)

func main() {

	var map1 map[string]any // declaration not instantiation

	if map1 == nil {
		println("map1 is nil")
	}

	//var map2 map[string]any = map[string]any{"560086": "Banglore-1", "560034": "Bangalore-2"}

	map1 = make(map[string]any, 20) //
	//println((map1))
	//cap(map1)

	map1["560086"] = "Bangalore-1"
	map1["560034"] = "Bangalore-3"
	map1["695011"] = "Trivandrum-1"
	map1["522001"] = "Guntur-1"
	map1["val"] = 100
	map1["done"] = true

	for key, value := range map1 {
		fmt.Println("key:", key, "value:", value)
	}

	v := map1["val"] // get a value from map
	fmt.Println(v)

	if v, ok := map1["val1"]; ok { // get a value from map
		fmt.Println(v)
	} else {
		println("key does not exist")
	}

	delete(map1, "done") // if key exists it deletes, if key does not .. it does nothing

	if v, ok := map1["done"]; ok { // get a value from map
		fmt.Println(v)
	} else {
		println("done key does not exist")
	}
	delete(map1, "done")

	if err := Delete(map1, "done"); err != nil {
		println(err.Error())
	} else {
		println("key successfully deleted")
	}

	if err := Delete(map1, "560086"); err != nil {
		println(err.Error())
	} else {
		println("key successfully deleted")
	}

	// slice1 := []int{10, 20}
	// slice2 := []int{10, 20}

	// if slice1 == slice2 {

	// }

	// arr1 := [2]int{10, 20}
	// arr2 := [2]int{10, 30}

	// if arr1 == arr2 {
	// 	println("both are same")
	// } else {
	// 	println("different")
	// }

	// map3 := make(map[[2]int]any)

	keys, values, err := GetKeysNValues(map1)

	if err != nil {
		println(err.Error())
	} else {
		fmt.Println("keys:", keys)
		fmt.Println("Values:", values)
	}

	clear(map1) // it clears the map all the keys and values are cleared from the map

	fmt.Println(map1) // clear does not make the map /slice nil but clears data in the map and make all the elements of the slice as zero values

	// slice fill numbers and check how many duplicate numbers are there

	slice := make([]int, 100)
	FillSlice(slice)

	fmt.Println(slice)

	mapSlice := make(map[int]int, len(slice))

	for _, v := range slice {
		val := mapSlice[v]
		mapSlice[v] = val + 1
	}

	for k, v := range mapSlice {
		fmt.Println("key:", k, "Value:", v)
	}

}

func FillSlice(slice []int) {
	for i := range slice {
		slice[i] = rand.IntN(1000)
	}
}

// make can be used for slice,maps and chan

// what type can be the key of a maq
// any type that can implement == operator

func Delete(m map[string]any, key string) error {
	if m == nil {
		return errors.New("nil map")
	}
	_, ok := m[key]
	if !ok {
		return errors.New("key:" + key + "->does not exist")
	}

	delete(m, key)
	return nil
}

func GetKeysNValues(m map[string]any) (keys []string, values []any, err error) {
	if m == nil {
		return nil, nil, errors.New("nil map")
	}
	keys = make([]string, len(m))
	index := 0
	for key, value := range m {
		keys[index] = key
		values = append(values, value)
		index++
	}
	return keys, values, nil
}
