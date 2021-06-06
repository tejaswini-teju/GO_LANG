package main

import (
	"fmt"
)

func main() {
	var num1 int
	var num2 int
	var num3 int
	fmt.Print("Enter the numbers ")
	fmt.Scanln(&num1, &num2, &num3)

	age := num1 + num2
	fmt.Println(age)

	if num1 > num2 && num1 > num3 {
		fmt.Println("Number 1 is greater")
	} else if num2 > num3 && num2 > num1 {
		fmt.Println("Number 2  is greatest")
	} else {
		fmt.Println("Number 3 is bigger")
	}

}
