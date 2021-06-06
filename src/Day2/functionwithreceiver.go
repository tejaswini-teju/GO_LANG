package main

import "fmt"

type person struct {
	fname string
	lname string
}

func (p person) details() {
	fmt.Println(p)
}

func main() {
	p1 := person{"teja", "nadella"}
	p1.details()

}
