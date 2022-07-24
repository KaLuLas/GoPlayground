package chapter11

import (
	"fmt"
	"math"
)

type Square struct {
	side float32
}

type Circle struct {
	radius float32
}

type Shaper interface {
	Area() float32
}

func (square Square) Area() float32 {
	return square.side * square.side
}

func (circle *Circle) Area() float32 {
	return circle.radius * circle.radius * math.Pi
}

func CheckShaperType() {
	var shaper Shaper
	//shaper = Square{side: 5}  // -> chapter11.Square
	shaper = &Square{side: 5} // -> *chapter11.Square
	//shaper = new(Square) // -> *chapter11.Square

	if t, ok := shaper.(*Square); ok {
		fmt.Printf("the type of %#v is: %T\n", shaper, t)
	} else if t, ok := shaper.(Square); ok {
		fmt.Printf("the type of %#v is: %T\n", shaper, t)
	}
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
