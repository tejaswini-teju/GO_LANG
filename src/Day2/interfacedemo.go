package main

import "fmt"

//interface
type human interface {
	sayHello() string
}

type man struct {
	greeting string
}
type woman struct {
	greeting string
}

func (m man) sayHello() string {
	return m.greeting
}
func (w woman) sayHello() string {
	return w.greeting
}
func printGreeting(h human) {
	fmt.Println(h.sayHello())
}

func main() {
	sai := man{greeting: "Hi"}
	siri := woman{greeting: "Hello"}

	printGreeting(sai)
	printGreeting(siri)
}
