package main

import (
	"fmt"
	"strconv"
)

const x = "hello World"

func main() {
	var i int = 42
	fmt.Printf("%v ,%T \n", i, i)
	//int  to string
	j := strconv.Itoa(i)
	fmt.Printf("%v ,%T", j, j)

	//string to int
	intres, err := strconv.ParseInt("10", 2, 32)
	fmt.Println(intres, err)

	//fmt.Println(x)

}
