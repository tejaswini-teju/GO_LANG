package main

import "fmt"

func normal() {
	fmt.Println("Inside Normal function")
}

func paramfunc(y int, x int) int {

	return x + y
}
func main() {
	//	normal()
	res := paramfunc(20, 10)
	fmt.Println(res)
}
