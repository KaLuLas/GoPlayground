package main

import (
	"chapter4"
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

func main() {
	result := utils.Add(a, utils.Add(b, c))
	fmt.Printf("result is %v\n", result)
	chapter4.PrintConst()
	chapter4.PrintOS()
	chapter4.TestConvert()
	chapter4.RunStrings()
	chapter4.RunStrConv()
	chapter4.TestTimePack()
	chapter4.TestIntPointer()
	chapter4.TestStrPointer()
}
