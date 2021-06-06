package main

import "fmt"

func main() {
	//int array
	arr := []int{1, 2, 3, 4, 5, 6}
	//i gives the index value
	for i, data := range arr {
		fmt.Println(i, " ", data)
	}

	for _, data := range arr {
		fmt.Println(" ", data)
	}
	//String array
	name := []string{"teju", "kolli", "mahesh"}
	for _, data := range name {
		fmt.Println(" ", data)
	}

	var intarr [3]int
	intarr[0] = 1
	intarr[1] = 2
	intarr[2] = 3
	fmt.Println(intarr)
	
	grades := [3]int{20, 30, 40}
	for _, data := range grades {
		fmt.Printf(" %v", data)
	}

}
