package chapter7

import "fmt"

func AppendSlice() {
	//slice1 := []int{1, 2, 3}
	slice1 := make([]int, 5, 10)
	slice2 := append(slice1, 4, 5, 6, 7, 8, 9, 10, 11) // bigger than 10, changes in slice2 will not affect slice1
	fmt.Println(slice1)
	fmt.Println(slice2)
	fmt.Println(&slice1 == &slice2) // different slice
}

func InsertStringSlice(slice, insertion []string, index int) []string {
	result := make([]string, len(slice)+len(insertion))
	pivot := copy(result, slice[:index])
	pivot += copy(result[index:], insertion)
	copy(result[pivot:], slice[index:])
	return result
}

func RemoveStringSlice(slice []string, from int, end int) []string {
	if end < from {
		return slice
	}

	if from < 0 || end >= len(slice) {
		return slice
	}

	// approach 1
	//result := make([]string, len(slice)-(end-from+1))
	//at := copy(result, slice[:from])
	//copy(result[at:], slice[end+1:])

	// approach 2
	result := append(slice[:from], slice[end+1:]...)

	return result
}
