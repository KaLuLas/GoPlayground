package chapter4

import (
	"fmt"
	"strconv"
	"strings"
)

func RunStrConv() {
	var orig string = "114509"
	var an int
	var newS string

	fmt.Printf("The size of ints is: %d\n", strconv.IntSize)

	an, _ = strconv.Atoi(orig)
	fmt.Printf("The integer is: %d\n", an)
	an = an + 5
	newS = strconv.Itoa(an)
	fmt.Printf("The new string is: %s\n", newS)
}

func RunStrings() {
	sample := "SampleString"
	println(strings.HasPrefix(sample, "Sample"))
	println(strings.Contains(sample, "leStr"))
	println(strings.Index(sample, "leStr"))
	println(strings.Repeat("Hello", 5))
	println(strings.ToUpper(sample))
	println(strings.ToLower(sample))
	reader := strings.NewReader(sample)

	for i := 0; i < len(sample); i++ {
		ch, _ := reader.ReadByte()
		print(ch, " ")
	}

	println()
}
