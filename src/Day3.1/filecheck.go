package main

import (
	"fmt"
	"os"
)

func main() {
	file, _ := os.OpenFile("newfile.txt", os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0644)
	file.WriteString("\n All")
	fmt.Println("Done")
}
