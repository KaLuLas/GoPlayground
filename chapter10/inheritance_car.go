package chapter10

import "fmt"

type Engine interface {
	Start()
	Stop()
}

type BasicEngine struct {
}

func (engine BasicEngine) Start() {
	fmt.Println("BasicEngine Start!")
}

func (engine BasicEngine) Stop() {
	fmt.Println("BasicEngine Stop!")
}

type Car struct {
	Engine
	Name       string
	wheelCount int
}

type Mercedes struct {
	Car
}

func (car *Car) numberOfWheels() int {
	return car.wheelCount
}

func (mercedes Mercedes) sayHiToMerKel() {
	fmt.Println("say hi to MerKel!")
}

func (mercedes Mercedes) SayHI() {
	mercedes.sayHiToMerKel()
}

func Finalizer(car *Car) {
	fmt.Printf("%s is garbage collected \n", car.Name)
}
