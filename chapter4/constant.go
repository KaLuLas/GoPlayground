package chapter4

import (
	"fmt"
)

// implicit
const Pi = 3.14159

// explicit
const Title string = "Pi"

const Monday, Tuesday, Wednesday = 1, 2, 3

// PrintConst print const value Pi and Monday
func PrintConst() {
	fmt.Printf("Pi is %v \n", Pi)
	fmt.Printf("Monday is %v\n", Monday)
}
