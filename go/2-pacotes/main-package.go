package main

import (
	"fmt"
	"modulo/support"

	"github.com/badoux/checkmail"
)

func main() {
	fmt.Println("Main module working")
	support.Write()

	error := checkmail.ValidateFormat("edudias@gmail.com")
	fmt.Println(error)
}
