package chapter6

import (
	"fmt"
	"strings"
)

func RunAdder() {
	var f = Adder(0)
	fmt.Print(f(1), " -> ")
	fmt.Print(f(20), " -> ")
	fmt.Println(f(300))
}

func Adder(base int) func(int) int {
	x := base
	return func(delta int) int {
		x += delta
		return x
	}
}

func MakeAddSuffix(suffix string) func(string) string {
	return func(name string) string {
		if !strings.HasSuffix(name, suffix) {
			return name + suffix
		}
		return name
	}
}
