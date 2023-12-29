package main

import "fmt"

func main() {
	slice := make([]float32, 10, 11)

	slice = append(slice, 4)
	slice = append(slice, 434)

	fmt.Println(slice)
	fmt.Println(len(slice)) //length
	fmt.Println(cap(slice)) //capacidade
}
