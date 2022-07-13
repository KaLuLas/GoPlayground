package chapter10

import "fmt"

type BaseObject struct {
	in1 int
	in2 int
}

type DerivedObject struct {
	b   string
	c   string
	int // anonymous field with built-in type
	BaseObject
}

func NewDerivedObject() *DerivedObject {
	object := new(DerivedObject)
	object.in1 = 1
	object.in2 = 2
	object.int = 3
	object.b = "stringB"
	object.c = "stringC"

	object1 := &DerivedObject{"stringB", "stringC", 1, BaseObject{2, 3}}
	if object1 == object {
		fmt.Printf("It's a f**king miracle")
	}

	return object
}
