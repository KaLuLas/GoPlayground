package chapter11

import (
	"fmt"
	"math"
	"unsafe"
)

type Square struct {
	side float32
}

type Circle struct {
	radius float32
}

type Triangle struct {
	width  float32
	height float32
}

type Shaper interface {
	Area() float32
	Direction() float32
}

func (square Square) Area() float32 {
	return square.side * square.side
}

func (square *Square) Direction() float32 {
	return 1.0
}

func (circle *Circle) Area() float32 {
	return circle.radius * circle.radius * math.Pi
}

func (circle *Circle) Direction() float32 {
	return -1.0
}

func (t Triangle) Area() float32 {
	return t.width * t.height / 2
}

func (t Triangle) Direction() float32 {
	//TODO implement me
	panic("implement me")
}

func CheckShaperType() {
	var shaper Shaper
	//shaper = Square{side: 5}  // -> chapter11.Square
	shaper = &Square{side: 5} // -> *chapter11.Square
	//shaper = new(Square) // -> *chapter11.Square

	if t, ok := shaper.(*Square); ok {
		fmt.Printf("the type of %#v is: %T\n", shaper, t)
	} /* else if t, ok := shaper.(Square); ok { // <- compile error
		fmt.Printf("the type of %#v is: %T\n", shaper, t)
	} */

	var shaper1 Shaper = Triangle{3, 4}
	fmt.Println("Triangle Shaper with size:", unsafe.Sizeof(shaper1)) // 16
	fmt.Println("Square Shaper with size:", unsafe.Sizeof(shaper))    // 16
}

func CheckShaperWithSwitch() {
	var shaper Shaper = &Square{side: 5}

	switch shaper.(type) {
	case *Square:
		fmt.Printf("the varaible %#v is *Square type", shaper)
	case *Circle:
		fmt.Printf("the varaible %#v is *Circle type", shaper)
	}
}
