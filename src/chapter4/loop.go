package chapter4

import "fmt"

func PrintLoop() {
	for i := 0; i < 5; i++ {
		fmt.Printf("This is the %d iteration\n", i)
	}
	// 无限循环
	for {
		fmt.Printf("looping...")
	}
}

func PrintWhileLoop() {
	i := 5
	// 基于条件判断的迭代，类似while结构
	for i >= 0 {
		fmt.Printf("This is the %d iteration\n", i)
		i -= 1
	}
}

func PrintTextIteration() {
	str := "一段标准中文文本"
	// 根据utf8字符迭代
	for pos, char := range str {
		fmt.Printf("character on position %d is %c\n", pos, char)
	}
}
