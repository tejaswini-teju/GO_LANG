package main

import (
	"fmt"
	"math/rand"
)

func main() {
	fmt.Println(rand.Intn(70))
	fmt.Println(rand.Float64())
	fmt.Println(rand.Intn(20))

	//perm function in the rand package gives pseudo random numbers
	v := rand.Perm(20)
	fmt.Println(v)

}
