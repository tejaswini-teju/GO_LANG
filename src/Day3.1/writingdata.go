package main

import (
	"fmt"
	"os"
)

func main() {
	f, _ := os.Create("Data.txt")
	fmt.Println("File created")
	bdata := []byte("Hi..Welcome to my world.")
	f.Write(bdata)

	//another way to insert data into file
	f.WriteString("My Name is Teju")
	fmt.Println("Data added")
	f.Close()
}
