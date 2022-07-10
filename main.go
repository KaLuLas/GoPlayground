package main

import (
	"GoPlayground/chapter10"
	"GoPlayground/chapter4"
	"GoPlayground/chapter6"
	"GoPlayground/chapter7"
	"GoPlayground/chapter8"
	"GoPlayground/chapter9"
	"GoPlayground/utils"
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
	chapter8.MakeMaps()
	chapter8.CheckAndDeleteMap()
	chapter8.SortMap()
}

func runChapter9() {
	chapter9.SizeOfInt()
}

func runChapter10() {
	chapter10.PersonCreation()
	var mat = chapter10.NewMatrix(10, 10)
	fmt.Printf("new matrix %+v, with size %d", mat, chapter10.SizeOfMatrix())
}

func main() {
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
}
