package main

import "fmt"

func main() {

	a := []string{"sony", "scooty"}
	for i, s := range a {
		fmt.Println(i, s)
	}

	mylist := []string{"java", "c", "go"}
	for _, lang := range mylist {
		fmt.Println(lang)
	}

}
