package main

import "fmt"

func main() {
	fmt.Println("Ponteiros")

	var numberOne int = 10

	var numberTwo int = numberOne

	numberOne++
	fmt.Println(numberOne, numberTwo)

	//PONTEIRO: REFERENCIA DE MEMORIA
	var numberThree int = 100
	var pointer *int = &numberThree

	numberThree += 103

	fmt.Println(numberThree, *pointer) //dereferencing
}
