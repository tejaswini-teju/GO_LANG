package main

import "fmt"

func main() {
	// func(a int, b int) {
	// 	fmt.Println(a + b)
	// }(1, 2)

	sum := func(a int, b int) {
		fmt.Println(a + b)
	}
	sum(15, 20)
}
