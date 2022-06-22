package chapter7

func ChangeFirstCharInString(input string, value byte) string {
	charArray := []byte(input)
	charArray[0] = value
	return string(charArray)
}

func ReverseString(input string) string {
	c := []byte(input)
	var tmp byte
	for i := 0; i < len(c)/2; i++ {
		other := len(c) - i - 1
		tmp = c[i]
		c[i] = c[other]
		c[other] = tmp
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
