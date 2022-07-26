package chapter11

import (
	"GoPlayground/utils/math"
	"fmt"
	"unsafe"
)

func PrintEmptyInterfaceSize() {
	var val interface{} = 1
	fmt.Println("empty interface with int has size:", unsafe.Sizeof(val))
	val = float64(3.14)
	fmt.Println("empty interface with float64 has size:", unsafe.Sizeof(val))
	val = math.NewVector3(3, 4, 5)
	fmt.Println("empty interface with Vector3 has size:", unsafe.Sizeof(val))
}
