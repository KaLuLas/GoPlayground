package utils

import (
	"fmt"
)

var utilsLabel = "edwardchen"

func init() {
	fmt.Printf("label is %v\n", utilsLabel)
	fmt.Println("utils.init executed.")
}

// this is a private add method
func add(param1, param2 int) int {
	fmt.Println("this is from utils.Add")
	return param1 + param2
}

// Add this is a public add method
func Add(param1, param2 int) int {
	fmt.Printf("add up %v and %v \n", param1, param2)
	return param1 + param2
}
