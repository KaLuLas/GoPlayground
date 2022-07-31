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

func EmptyInterface() {
	type Any interface{}
	var val Any
	val = 5
	fmt.Printf("val has the value: %v\n", val)
	val = "SampleString"
	fmt.Printf("val has the value: %v\n", val)
	val = math.NewVector3(3, 4, 5)
	fmt.Printf("val has the value: %v\n", val)
	switch t := val.(type) {
	case int:
		fmt.Printf("val has type %T\n", t)
	case string:
		fmt.Printf("val has type %T\n", t)
	case *math.Vector3:
		fmt.Printf("val has type %T\n", t)
	default:
		fmt.Printf("unexpected type %T\n", t)
	}
}
