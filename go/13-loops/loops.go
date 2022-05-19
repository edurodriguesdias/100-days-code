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

	for _, value := range names {
		fmt.Println(value)
	}

	for index, letter := range "WORD" {
		fmt.Println(index, string(letter))
	}

	user := map[string]string{
		"name":    "João",
		"surname": "Silva",
	}

	for index, value := range user {
		fmt.Println(index, value)
	}
}
