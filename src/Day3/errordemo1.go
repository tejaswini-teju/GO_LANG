package main

import "fmt"

func main() {
	fmt.Println(Divide(100, 0))
	fmt.Println(Divide(100, 10))
	fmt.Println(Divide(20, 0))
}

func Divide(num1, num2 int) int {
	defer func() {
		fmt.Println(recover())
	}()
	quotient := num1 / num2
	return quotient
}
