package chapter7

import (
	"fmt"
	"utils"
)

func MakeIntSlice(length int, cap int, value int) {
	// initiate with 'make'
	slice1 := make([]int, length, cap)
	for i := range slice1 {
		slice1[i] = value
	}
}

func GrowSlice() {
	slice1 := make([]int, 0, 5)
	// grow length of slice from 0 to 5
	for i := 0; i < cap(slice1); i++ { // notice: cap(slice1)
		slice1 = slice1[0 : i+1]
		slice1[i] = i
		fmt.Printf("length of slice is now %d \n", len(slice1))
	}

	utils.LogSlice(slice1, "grow_slice_example")
}

func Reslice() {
	array := [10]int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	slice1 := array[5:7]
	utils.LogSlice(slice1, "reslice_example")
	// actually print slice is good enough, get "[5 6]"
	fmt.Println(slice1)

	slice1 = slice1[0:4]
	utils.LogSlice(slice1, "reslice_example")
	fmt.Println(slice1)
	// not available
	//slice1 = slice1[-1:4]
}
