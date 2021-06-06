package main

import (
	"fmt"
	"strings"
)

func main() {
	info := "My name is Teju"
	lang := "golang,java,c,sql"
	fmt.Println(strings.Compare(info, "My name is teju"))
	fmt.Println(strings.Contains(info, "Teju"))

	fmt.Println(strings.Count(info, "a"))
	fmt.Println(len(info))
	fmt.Println(strings.Index(info, "n")) //starts from 0th
	fmt.Println(strings.Split(info, " "))

	//to count no of substrings
	fmt.Println(strings.SplitN(lang, ",", 3))
	var arr []string
	arr = strings.SplitN(lang, ",", 3)
	fmt.Println(len(arr))

	//strings are immutable
	fmt.Println(strings.ToUpper(lang))
	fmt.Println(lang)

	a := info
	b := a[3:5]
	print(b)

	//string to int

}
