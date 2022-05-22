package chapter6

import (
	"log"
	"runtime"
)

var Where = func() {
	_, file, line, _ := runtime.Caller(1)
	log.Printf("caller:%s line:%d", file, line)
}
