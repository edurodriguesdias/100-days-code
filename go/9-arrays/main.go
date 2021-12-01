package main

import (
	"fmt"
	"reflect"
)

func main() {

	var array1 [5]int

	array2 := [3]string{"um", "dois", "tres"}

	array3 := [...]int{1, 2, 3, 4, 5, 6, 7}

	fmt.Println(array1, array2, array3)

	fmt.Println(reflect.TypeOf(array3))
}
