package main

import "fmt"

func main() {
	users := map[string]string{
		"name":    "João",
		"surname": "Silva",
	}

	delete(users, "name")
	fmt.Println(users)
}
