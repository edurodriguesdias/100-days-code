package main

import (
	"fmt"
)

func main() {
	fmt.Println("Estrutura de controle")

	number := 10

	if number > 15 {
		fmt.Println("Maior que 15")
	} else {
		fmt.Println("Menor que 15")
	}

	if otherNumber := number; otherNumber > 0 { //Var limitada ao escopo do if
		fmt.Println("Mair que zero")
	}

}
