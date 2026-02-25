package main

import (
	"errors"
	"fmt"
)

type MyAnotherMap = map[string]any // typeDef , it is not a new type just a name to another type

type MyAnotherMap1 map[string]any

func main() {

	m1 := make(MyMap)

	m1["name"] = "Jiten"
	m1["email"] = "jitenp@outlook.com"
	m1["mobile"] = 9618558500
	m1["done"] = true

	if err := m1.Delete("done"); err != nil {
		println(err.Error())
	} else {
		println("key/value pair successfully deleted")
	}

	if keys, values, err := m1.GetKeysNValues(); err != nil {
		println(err.Error())
	} else {
		fmt.Println("Keys:", keys, "\nvalues:", values)
	}

	m2 := make(map[string]any)

	m2["name"] = "Jiten"
	m2["email"] = "jitenp@outlook.com"
	m2["mobile"] = 9618558500
	m2["done"] = true

	if err := MyMap(m2).Delete("done"); err != nil {
		println(err.Error())
	} else {
		println("key/value pair successfully deleted")
	}

	if keys, values, err := MyMap(m2).GetKeysNValues(); err != nil {
		println(err.Error())
	} else {
		fmt.Println("Keys:", keys, "\nvalues:", values)
	}

	m3 := make(MyAnotherMap)

	m3["name"] = "Jiten"
	m3["email"] = "jitenp@outlook.com"
	m3["mobile"] = 9618558500
	m3["done"] = true

	slice := make([]any, 3)
	slice[0] = m1
	slice[1] = m2
	slice[2] = m3

	for _, v := range slice {
		switch m := v.(type) {

		case MyMap:
			println("\nMyMap")

			if keys, values, err := m.GetKeysNValues(); err != nil {
				//if keys, values, err := v.(MyMap).GetKeysNValues(); err != nil {
				println(err.Error())
			} else {
				fmt.Println("Keys:", keys, "\nvalues:", values)
			}
		case map[string]any:
			println("\nmap[string]any")
			if keys, values, err := MyMap(m).GetKeysNValues(); err != nil {
				//if keys, values, err := MyMap(v.(map[string]any)).GetKeysNValues(); err != nil {
				println(err.Error())
			} else {
				fmt.Println("Keys:", keys, "\nvalues:", values)
			}
		default:
			println("unidentified type")
		}

	}

}

type MyMap map[string]any

func (m MyMap) Delete(key string) error {

	if m == nil {
		return errors.New("nil map")
	}

	_, ok := m[key]
	if ok {
		delete(m, key)
	} else {
		return errors.New("key >>" + key + " does not exist")
	}

	return nil
}

func (m MyMap) GetKeysNValues() (keys []string, values []any, err error) {
	if m == nil {
		return nil, nil, errors.New("nil map")
	}
	for k, v := range m {
		keys = append(keys, k)
		values = append(values, v)
	}
	return keys, values, nil
}
