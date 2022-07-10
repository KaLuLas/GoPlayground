package chapter10

import "unsafe"

type matrix struct {
	row int
	col int
}

func NewMatrix(row int, col int) *matrix {
	m := &matrix{row, col}
	return m
}

func SizeOfMatrix() uintptr {
	return unsafe.Sizeof(matrix{})
}
