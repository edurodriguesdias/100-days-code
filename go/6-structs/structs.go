package main

import "fmt"

type user struct {
	name    string
	age     byte
	address address
}

type address struct {
	street string
	number byte
}

func main() {
	var customer user

	customer.name = "Eduardo"
	customer.age = 25
	fmt.Println(customer)

	addressExample := address{"Rua pedro s", 202}

	seller := user{"João", 30, addressExample}
	fmt.Println(seller)

	driver := user{age: 21}
	fmt.Println(driver)
}
