package chapter6

import "fmt"

func GetX2AndX3(input int) (int, int) {
	return 2 * input, 3 * input
}

func GetX2AndX3_2(input int) (x2 int, x3 int) {
	defer func() {
		fmt.Println("x2 is", x2)
		fmt.Println("x3 is", x3)
	}()
	x2 = input * 2
	x3 = input * 3
	//return
	//return x2, x3
	return input, input
}

func PrintStackTrace() {
	b()
}

func b() {
	defer un(trace("b"))
	fmt.Println("in b")
	a()
}

func a() {
	defer un(trace("a"))
	fmt.Println("in a")
}

func un(s string) {
	fmt.Println("leaving:", s)
}

func trace(s string) string {
	fmt.Println("entering:", s)
	return s
}
