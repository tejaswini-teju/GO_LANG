package main

import (
	"fmt"
	"reflect"
)

func main() {
	num1 := 10
	fmt.Println(num1)
	fmt.Println(&num1)

	num2 := &num1
	fmt.Println(num2)
	fmt.Println(reflect.TypeOf(num2))

	//change the value of num2
	*num2 = 10
	fmt.Println(num1)
	fmt.Println(*num2)

	// a := 42
	// fmt.Println(a)
	// fmt.Println(&a)

	// //b stores address of a b =&a
	// var b *int = &a
	// fmt.Println(a)
	// fmt.Println(b)
	// //gives values of a where a = *b
	// fmt.Println(*b)

}
