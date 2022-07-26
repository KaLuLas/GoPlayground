package chapter7

func ChangeFirstCharInString(input string, value byte) string {
	charArray := []byte(input)
	charArray[0] = value
	return string(charArray)
}

func ReverseString(input string) string {
	c := []byte(input)
	for i := 0; i < len(c)/2; i++ {
		other := len(c) - i - 1
		c[i], c[other] = c[other], c[i]
	}

	return string(c)
}

func MapFunction(processor func(int) int, input []int) []int {
	output := make([]int, len(input))
	for i := 0; i < len(input); i++ {
		output[i] = processor(input[i])
	}

	return output
}
