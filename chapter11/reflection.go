package chapter11

import (
	"GoPlayground/chapter10"
	"fmt"
	"reflect"
)

func ReflectFloat() {
	x := 3.14
	v := reflect.ValueOf(x)
	fmt.Println("Value of x is", v)
	fmt.Println("Type of v is", v.Type())
	fmt.Println("Kind if v is", v.Kind())
	fmt.Println("Settability of v is", v.CanSet())
	fmt.Println("'CanFloat' of v is", v.CanFloat())
	fmt.Println("float state of v is", v.Float())
}

func ReflectCar() {
	car := new(chapter10.Car)
	car.Engine = new(chapter10.BasicEngine)
	car.Name = "TestCarForReflection"
	v := reflect.ValueOf(car).Elem()
	t := v.Type()
	for i := 0; i < v.NumField(); i++ {
		field := v.Field(i)
		fmt.Printf("Field %d: %s %s \n", i, t.Field(i).Name, field.Type())
	}

	for i := 0; i < v.NumMethod(); i++ {
		method := v.Method(i)
		fmt.Printf("Method %d: %s %s \n", i, t.Field(i).Name, method.Type())
	}
}
