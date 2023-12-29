package main

import "fmt"

type User struct {
	name    string
	age     byte
	address address
}

type Address struct {
	street string
	number byte
}

func main() {
	var customer User

	customer.name = "Eduardo"
	customer.age = 25
	fmt.Println(customer)

	addressExample := Address{"Rua pedro s", 202}

	seller := User{"João", 30, addressExample}
	fmt.Println(seller)

	driver := User{age: 21}
	fmt.Println(driver)
}
