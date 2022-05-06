package chapter4

import (
	"fmt"
	"os"
	"runtime"
)

func PrintOS() {
	path := os.Getenv("GOPATH")
	fmt.Printf("The operating system is %s:%s\n", runtime.GOOS, runtime.GOARCH)
	fmt.Printf("GOPATH is %s\n", path)
}
