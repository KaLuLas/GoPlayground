package utils

import "fmt"

func LogSlice(slice []int, identifier string) {
	if slice == nil {
		return
	}

	fmt.Printf("length of %s is %d, cap of slice is %d \n", identifier, len(slice), cap(slice))
	for index, value := range slice {
		fmt.Printf("element at %s no.%d is %d \n", identifier, index, value)
	}
}
