package chapter4

import "fmt"

func TestIntPointer() {
	i1 := 5
	i2 := 6
	fmt.Printf("An integer: %d, its location in memory: %p\n", i1, &i1)
	fmt.Printf("An integer: %d, its location in memory: %p\n", i2, &i2)
	intPtr := &i1
	fmt.Printf("The value at memory location %p is %d\n", intPtr, *intPtr)
}

func TestStrPointer() {
	s := "good bye"
	var strPtr *string = &s
	fmt.Printf("previous 'strPtr' location is %p\n", strPtr)
	*strPtr = "hello"
	fmt.Printf("now 'strPtr' location is %p\n", strPtr)
	fmt.Printf("now string 's' is %s\n", s)
}
