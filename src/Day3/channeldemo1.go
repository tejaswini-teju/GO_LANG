package main

import (
	"fmt"
	"runtime"
	"time"
)

func main() {
	n := 3
	//run go routine to multiply n by 2
	//make the channge;
	out := make(chan int)
	go multiplyByTwo(n, out)
	time.Sleep(time.Second)

	fmt.Println(<-out)

}
func multiplyByTwo(num int, out chan<- int) {
	result := num * 2
	//fmt.Println(result)
	out <- result
	//return result
	fmt.Println(runtime.)
}
