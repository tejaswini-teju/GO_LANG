package main

import (
	"fmt"
	"os"
)

func main() {
	fd := os.Mkdir("sample", 0644)
	fmt.Println("folder created", fd)
}

