package main

import (
	"fmt"
	"os"
)

func main() {
	fp, _ := os.Create("newfile.txt")
	fmt.Println("file created", fp)
}
