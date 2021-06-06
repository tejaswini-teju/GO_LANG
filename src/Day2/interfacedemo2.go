package main

import "fmt"

type vehicle interface {
	engine() string
}

type audi struct {
	model string
	name  string
}

type polo struct {
	model string
	name  string
}

func vdetails(v vehicle) {
	fmt.Println(v)
}

func (a audi) engine() string {
	fmt.Println("Audi is auto engine")
	return a.model
}

func (p polo) engine() string {
	fmt.Println("polo is manual")
	return p.model
}

func main() {
	audiObj := audi{"A8", "audi"}
	audiObj.engine()
	poloobj := polo{"2019", "polo"}
	poloobj.engine()
	vdetails(poloobj)
}
