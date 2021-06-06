package main

import (
	"fmt"
	"runtime"
)

func main() {
	fmt.Println(runtime.GOOS)     //os details
	fmt.Println(runtime.NumCPU()) // gives no of cores

	fmt.Println(runtime.NumGoroutine())
}
