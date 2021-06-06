package main

import (
	"fmt"
	"reflect"
)

var name = "Sony"

func main() {
	num := 20
	fmt.Printf("num is of type %T \n", num)
	fmt.Print(reflect.TypeOf(num))
	
}
