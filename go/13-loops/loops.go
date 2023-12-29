package main

import (
	"fmt"
	"time"
)

func main() {
	increment := 0

	for increment <= 10 {
		time.Sleep(time.Millisecond)
		increment++
	}

	fmt.Println(increment)

	for another := 0; another <= 10; another++ {
		fmt.Println("increasing another", another)
		time.Sleep(time.Millisecond)
	}

	names := [3]string{"Harry", "Kate", "Jessi"}

	for _, value := range names { //Iterar sobre uma lista
		fmt.Println(value)
	}

	for index, letter := range "WORD" { //Iterar sobre uma string, mas por padrão é retornado os códigos ASC
		fmt.Println(index, string(letter)) //uso da função string para "decodificar" os códigos ASC para letras
	}

	User := map[string]string{
		"name":    "João",
		"surname": "Silva",
	}

	for _, value := range User {
		fmt.Println(value)
	}
}
