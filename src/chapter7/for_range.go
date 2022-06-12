package chapter7

import (
	"fmt"
)

func MultiDimSlice() {
	//var screen [5][5]int
	screen1 := new([5][5]int) // with 'new'
	for row := range *screen1 {
		for col := range (*screen1)[row] {
			(*screen1)[row][col] = 1
			fmt.Printf("row:%d col:%d value %d \n", row, col, (*screen1)[row][col])
		}
	}

	// also available, how?
	//for row := range screen1 {
	//	for col := range screen1[row] {
	//		screen1[row][col] = 1
	//		fmt.Printf("row:%d col:%d value %d \n", row, col, screen1[row][col])
	//	}
	//}

	var screen [5][5]int // without 'new'
	for row := range screen {
		for col := range screen[row] {
			if row >= col {
				screen[row][col] = row
			} else {
				screen[row][col] = col
			}

			fmt.Printf("row:%d col:%d value %d \n", row, col, screen1[row][col])
		}
	}
}
