package main

import "fmt"

func add(y int, x int) int {
	return x + y
}

func multi(x int, y int) (int, int) {
	return x * y, x / y
}
func main() {

	res := add(20, 10)
	fmt.Println(res)
	multi, div := multi(12, 4)
	fmt.Println("Multiplication result :", multi, "Division Result :", div)
}
