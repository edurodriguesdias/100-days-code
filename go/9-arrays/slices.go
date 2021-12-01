package main

import (
	"fmt"
	"reflect"
)

func main() {
	slice := []int{19, 12, 23, 3123, 312, 3123, 23, 123, 23, 123, 23, 123}
	array := [4]string{"Zero", "Um", "Dois", "Três"}

	fmt.Println(reflect.TypeOf(slice))

	slice = append(slice, 1235)

	slice2 := array[1:3]

	fmt.Println(slice, slice2)
}
