package chapter4

import (
	"fmt"
	"math"
)

func Uint8FromInt(n int) uint8 {
	if 0 <= n && n <= math.MaxUint8 {
		return uint8(n)
	}

	panic(fmt.Sprintf("%d is out of the uint8 range", n))
}

func IntFromFloat64(x float64) int {
	if math.MinInt32 <= x && x <= math.MaxInt32 {
		whole, frac := math.Modf(x)
		if frac >= 0.5 {
			whole++
		}

		return int(whole)
	}

	panic(fmt.Sprintf("%g is out of the int32 range", x))
}

func TestConvert() {
	result := Uint8FromInt(255)
	fmt.Printf("uint8 %v from int\n", result)
	result1 := IntFromFloat64(1.14514)
	fmt.Printf("int %v from float64\n", result1)
}
