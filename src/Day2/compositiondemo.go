package main

import "fmt"

type address struct {
	pin  int
	city string
}

type student struct {
	name string
	id   int
	age  int
	sadd address
}

func main() {
	//create address struct and pass it to student
	addobj := address{521135, "vijayawada"}
	//pass address type to student
	sobj := student{"siri", 101, 20, addobj}
	fmt.Println(sobj)
	fmt.Println(sobj.sadd, sobj.name)
	//key value pair
	s1obj := student{name: "sai", id: 102, age: 20, sadd: address{500021, "vvnagar"}}
	fmt.Println(s1obj.id)
	fmt.Println(s1obj.sadd.pin)

}
