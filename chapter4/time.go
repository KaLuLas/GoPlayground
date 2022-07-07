package chapter4

import (
	"fmt"
	"time"
)

func TestTimePack() {
	t := time.Now()
	fmt.Println(t)
	fmt.Println(t.UTC())
	fmt.Println(t.Format(time.RFC3339))
}
