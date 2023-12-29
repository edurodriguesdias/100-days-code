package main

import "fmt"

type Person struct {
	name    string
	surname string
	age     uint8
	height  uint8
}

type Student struct {
	Person
	course  string
	college string
}

func main() {
	fmt.Println("Inheritance")

	people := Person{"Eduardo", "Dias", 25, 175}

	peopleStudent := Student{people, "Analise de sistemas", "USP"}

	fmt.Println(people)
	fmt.Println(peopleStudent)

	fmt.Println(peopleStudent.name)
	fmt.Println(peopleStudent.height)
}
