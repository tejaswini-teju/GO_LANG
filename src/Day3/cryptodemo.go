package main

import (
	"fmt"

	"golang.org/x/crypto/bcrypt"
)

func main() {
	pass := "admin"
	newpwd, _ := bcrypt.GenerateFromPassword([]byte(pass), 4)
	fmt.Println(newpwd)
	//decryption
	err := bcrypt.CompareHashAndPassword(newpwd, []byte("admin"))
	if err != nil {
		fmt.Println("Invalid password")
	} else {
		fmt.Println("welcome user password matched")
	}
}
