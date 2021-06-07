package main

import "fmt"

const (
	a = iota
	b
	c
)

//enumns can be created using iota

type Size int

const (
	length = Size(iota + 10)
	breadth
	height
)

func (s Size) toString() {
	switch s {
	case length:
		fmt.Println("Small")
	case breadth:
		fmt.Println("Medium")
	case height:
		fmt.Println("Extra Large")
	default:
		fmt.Println("Invalid Size entry")
	}
}

func main() {
	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)
	var m Size = 10
	m.toString()
	fmt.Println(length)
	fmt.Println(breadth)
	fmt.Println(height)
}
