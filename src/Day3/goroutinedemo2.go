package main

import (
	"fmt"
	"sync"
	"time"
)

var wg = sync.WaitGroup{}

func sayhello() {
	for i := 0; i < 5; i++ {
		fmt.Println("Good morning")
		time.Sleep(1000 * time.Millisecond)
	}
	wg.Done()
}
func saybye() {
	for i := 0; i < 5; i++ {
		fmt.Println("Good day")
		time.Sleep(1000 * time.Millisecond)
	}
	wg.Done()
}
func main() {
	wg.Add(2) //no of go rotines
	//waitgroup willwait for 2 go routines
	go sayhello()
	go saybye()
	wg.Wait() //wait till goroutine give acknowledgement "done"
	fmt.Println("Last line in main")
}
