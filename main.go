package main

import (
	"GoPlayground/chapter10"
	"GoPlayground/chapter11"
	"GoPlayground/chapter4"
	"GoPlayground/chapter6"
	"GoPlayground/chapter7"
	"GoPlayground/chapter8"
	"GoPlayground/chapter9"
	"GoPlayground/utils"
	"GoPlayground/utils/math"
	"fmt"
)

var a int     // default value
var b int = 1 // type explicit
var c = 2     // type implicit

// init method example
func init() {
	fmt.Println("main.init executed.")
}

func loadUtils() {
	result := utils.Add(a, utils.Add(b, c))
	fmt.Printf("result is %v\n", result)
}

func runChapter7() {
	fmt.Println("///// CHAPTER 7 /////")
	chapter7.SliceExample()
	//chapter7.MultiDimSlice()
	//chapter7.GrowSlice()
	//chapter7.Reslice()
	chapter7.AppendSlice()
	result := chapter7.InsertStringSlice([]string{"a", "b", "c"}, []string{"e", "f", "g"}, 1)
	fmt.Println(result)
	result = chapter7.RemoveStringSlice([]string{"a", "b", "c", "d"}, 0, 2)
	fmt.Println(result)
	fmt.Println(chapter7.ChangeFirstCharInString("iimple", 's'))
	fmt.Println(chapter7.ReverseString("Kalulas"))
	chapter7.MapFunction(func(input int) int {
		return input * 10
	}, []int{1, 2, 3, 4})
}

func runChapter8() {
	fmt.Println("///// CHAPTER 8 /////")
	chapter8.MakeMaps()
	chapter8.CheckAndDeleteMap()
	chapter8.SortMap()
}

func runChapter9() {
	fmt.Println("///// CHAPTER 9 /////")
	chapter9.SizeOfInt()
}

func runChapter10() {
	fmt.Println("///// CHAPTER 10 /////")
	chapter10.PersonCreation()
	var mat = chapter10.NewMatrix(10, 10)
	fmt.Printf("new matrix %+v, with size %d\n", mat, chapter10.SizeOfMatrix())
	dObject := chapter10.NewDerivedObject()
	fmt.Printf("%+v \n", dObject)
	// combination
	car := &chapter10.Mercedes{}
	car.Name = "MyFirstCar"
	//car = new(chapter10.Mercedes) // the same
	car.SayHI()
	//runtime.SetFinalizer(car, chapter10.Finalizer)
	// inheritance
	var engine chapter10.Engine
	engine = new(chapter10.BasicEngine)
	engine.Start()

	engineSlice := make([]chapter10.Engine, 1)
	engineSlice = append(engineSlice, engine)
}

func runChapter11() {
	fmt.Println("///// CHAPTER 11 /////")
	chapter11.CheckShaperType()
	chapter11.CheckShaperWithSwitch()
	chapter11.TrySortDays()
	chapter11.TrySortPersons()
	chapter11.PrintEmptyInterfaceSize()
	chapter11.EmptyInterface()
	chapter11.ReflectFloat()
	chapter11.ReflectCar()
}

func runUtilsMath() {
	vec3 := math.NewVector3(3.0, 4.0, 5.0)
	fmt.Printf("%+v with abs %f \n", vec3, vec3.Abs())
}

func main() {
	runUtilsMath()
	//loadUtils()
	//chapter4.PrintConst()
	chapter4.PrintOS()
	//chapter4.TestConvert()
	//chapter4.RunStrings()
	//chapter4.RunStrConv()
	chapter4.TestTimePack()
	//chapter4.TestIntPointer()
	//chapter4.TestStrPointer()
	chapter4.PrintTextIteration()
	//chapter4.PrintGotoLabel()
	ret1, ret2 := chapter6.GetX2AndX3_2(10)
	fmt.Printf("ret1 is %d, ret2 is %d\n", ret1, ret2)
	//chapter6.PrintStackTrace()
	chapter6.RunAdder()
	aviAdder := chapter6.MakeAddSuffix(".avi")
	fmt.Println(aviAdder("chapter4"))
	//chapter6.Where()
	runChapter7()
	runChapter8()
	runChapter9()
	runChapter10()
	runChapter11()
}
