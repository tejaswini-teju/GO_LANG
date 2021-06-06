package main

import (
	"fmt"
)

func main() {
	var name string
	var fname string

	fmt.Print("Enter your name:")
	fmt.Scanln(&name)
	fmt.Println(name)

	fmt.Scanf("%q", &fname)
	fmt.Println(fname)

}
