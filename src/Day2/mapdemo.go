package main

import "fmt"

func main() {
	mymap := map[int]string{1: "a", 2: "b", 3: "c"}
	mymap[10] = "d"
	fmt.Println(mymap)

	map2 := map[int]int{1: 20, 2: 30, 3: 40}
	fmt.Println(map2)

	delete(mymap, 10)
	fmt.Println(mymap)
}
