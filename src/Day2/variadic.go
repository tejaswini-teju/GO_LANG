package main

import "fmt"

func add(args ...int) {
	fmt.Print(args)
	total := 0
	for _, v := range args {
		total += v
	}
	fmt.Println(total)
}
func main() {
	add(10, 20)
	nums := []int{1, 2, 3}
	add(nums...)
}
