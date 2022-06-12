package chapter7

import "fmt"

var emptyArray [5]int      // empty array
var emptyArray1 = [5]int{} // empty array, '{}' is necessary

func SliceExample() {
	var arr1 [6]int
	//var slice1 []int = arr1[2:5]
	slice1 := arr1[2:5] // 1. from array
	//slice1 := make([]int, 0, 5) // 2. with 'make'
	//slice1 := []int{1,2,3,4,5} // 3. without 'make'

	for i := 0; i < len(arr1); i++ {
		arr1[i] = i
	}

	for i := 0; i < len(slice1); i++ {
		fmt.Printf("Slice at %d is %d \n", i, slice1[i])
	}

	fmt.Printf("length of array is %d \n", len(arr1))
	fmt.Printf("length of slice is %d \n", len(slice1))
	fmt.Printf("capacity of slice is %d \n", cap(slice1))

	//slice1 = slice1[0:4]
	slice1 = arr1[0:4]
	for i := 0; i < len(slice1); i++ {
		fmt.Printf("Slice at %d is %d \n", i, slice1[i])
	}

	fmt.Printf("length of slice is %d \n", len(slice1))
	fmt.Printf("capacity of slice is %d \n", cap(slice1))
}
