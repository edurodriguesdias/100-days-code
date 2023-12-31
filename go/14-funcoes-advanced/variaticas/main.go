package main

import "fmt"

func sum(nums ...int) (total int) {

	for _, num := range nums {
		total += num
	}

	return
}

func main() {
	resp := sum(1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 10, 10, 10, 200)

	fmt.Println(resp)
}
