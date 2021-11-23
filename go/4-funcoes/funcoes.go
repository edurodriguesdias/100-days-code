package main

import "fmt"

func sum(numberOne, numberTwo byte) byte {
	return numberOne + numberTwo
}

func mathCalcs(numberOne, numberTwo byte) (byte, byte) {
	sum := numberOne + numberTwo
	sub := numberOne - numberTwo

	return sum, sub
}

func main() {
	sum := sum(2, 33)
	fmt.Println(sum)

	var mathSum, _ byte = mathCalcs(20, 10)

	fmt.Println(mathSum)

	fmt.Println("jx")

	fmt.Println("jefferson")
}
