package main

import "fmt"

func main() {
	defer firstrun()
	secondrun()
}

func firstrun()  { fmt.Println("First run") }
func secondrun() { fmt.Println("second run") }
