package main

import (
	"fmt"
)

func main() {
	go hello()
	//time.Sleep(1 * time.Second)
	fmt.Println("Hi from main function")
}
func hello() {
	fmt.Println("Hello from go routine")
}
