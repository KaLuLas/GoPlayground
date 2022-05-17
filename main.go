package main

import (
	"chapter4"
	"chapter6"
	"fmt"
	"utils"
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
}
