package main

import (
	"fmt"
	"time"
)

func main() {
	ptime := time.Now()
	fmt.Println(ptime)
	time.Sleep(1000 * time.Microsecond)
	fmt.Println(ptime.Date())

	fmt.Println(ptime.Day())
	fmt.Println(ptime.Hour())
	fmt.Println(ptime.Year())

	fmt.Println(ptime.Local())
	fmt.Println(ptime.Clock())

	wday := time.Now().Weekday()
	fmt.Println(wday)
	fmt.Println(int(wday))

}
