package chapter9

import (
	"fmt"
	"unsafe"
)

func SizeOfInt() {
	i := 10
	size := unsafe.Sizeof(i)
	fmt.Println("The size of and int is: ", size)
}
