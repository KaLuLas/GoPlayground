package chapter7

import "fmt"

func AppendSlice() {
	//slice1 := []int{1, 2, 3}
	slice1 := make([]int, 5, 10)
	slice2 := append(slice1, 4, 5, 6)
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
	// TODO
	return []string{"to", "be", "continue"}
}
