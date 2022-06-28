package chapter8

import "fmt"

func CheckAndDeleteMap() {
	var value int
	var isPresent bool

	map1 := make(map[string]int)
	map1["keyA"] = 1
	map1["keyB"] = 2
	map1["keyC"] = 3
	value, isPresent = map1["keyA"]

	if isPresent {
		fmt.Printf("Value of 'keyA' is %d\n", value)
	}

	value, isPresent = map1["keyD"]
	fmt.Printf("Not existed key value %d, present %t \n", value, isPresent)

	delete(map1, "keyA")
	value, isPresent = map1["keyA"]
	fmt.Printf("'keyA' value %d, present %t \n", value, isPresent)
}
