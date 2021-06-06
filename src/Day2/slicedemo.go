package main

import "fmt"

func main() {
	a := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	fmt.Println(a)
	//prints entire array
	fmt.Println(a[:])
	//3rd to end
	fmt.Println(a[3:])
	//0 t0 5th element
	fmt.Println(a[:5])
	//slice 4,5,6
	fmt.Println(a[3:6])
}
