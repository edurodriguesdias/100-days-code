package main

import (
	"errors"
	"fmt"
)

func main() {
	var error = errors.New("error to create")

	fmt.Println(error)
}
