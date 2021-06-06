package main

import "fmt"

type employee struct {
	fname string
	lname string
	age   int
}

type rectangle struct {
	lenght  int
	breadth int
	color   string
}

func main() {
	//e := employee{fname: "Teja", lname: "Nadella", age: 16}
	e := employee{"teju", "kolli", 25}
	fmt.Println(e)
	fmt.Println(e.fname)

	//creating reference for struct using new
	rect := new(rectangle)
	rect.breadth = 10
	rect.lenght = 20
	rect.color = "blue"
	fmt.Println(rect)

	//struct initialisation using pointer
	var rect1 = &rectangle{}
	(*rect1).breadth = 10
	*&rect1.lenght = 20
	(*rect1).color = "white"
	fmt.Println(rect1)
}
