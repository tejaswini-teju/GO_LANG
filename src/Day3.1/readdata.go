package main

import (
	"fmt"
	"io/ioutil"
)

func main() {

	//to read file
	data, _ := ioutil.ReadFile("Data.txt")
	fmt.Println(string(data))
}
