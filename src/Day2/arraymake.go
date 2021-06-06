package main

import "fmt"

func main() {

	//make usage for map
	mymap := make(map[int]string)
	mymap[1] = "Teja"
	fmt.Println(mymap)

	//make usage for an integer list
	myslice := make([]int, 2, 3)
	fmt.Println(len(myslice))
	fmt.Println(cap(myslice))
	fmt.Println(myslice)

	myslice[0] = 15
	myslice[1] = 24
	fmt.Println(myslice)
	//using append
	myslice = append(myslice, 10)
	myslice = append(myslice, 45)
	myslice = append(myslice, 50)
	myslice = append(myslice, 60)
	myslice = append(myslice, 70)
	fmt.Println(myslice)
	fmt.Println(len(myslice))
	fmt.Println(cap(myslice))

	

}
