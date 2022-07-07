package chapter8

import "fmt"

func MakeMaps() {
	// declare
	//var map1 map[string]int

	// declare and make
	//var map1 map[string]int = make(map[string]int)

	// make with size
	//var map1 map[string]int = make(map[string]int, 10)

	//var map1 = make(map[string]int)
	mapCreated := make(map[string]int)
	mapLit := map[string]int{"one": 1, "two": 2}
	mapAssigned := mapCreated

	mapCreated["three"] = 3
	mapCreated["four"] = 4
	fmt.Println(mapCreated)
	mapAssigned["three"] = -3 // affect mapCreated
	mapAssigned = mapLit
	mapAssigned["one"] = -1 // affect mapLit
	fmt.Println(mapCreated)
	fmt.Println(mapLit)
}

func MapWithFunction() map[int]func() {
	return map[int]func(){}
}

func MapWithSlice() {

}
