package main

import "fmt"

type person struct {
	name    string
	surname string
	age     uint8
	height  uint8
}

type student struct {
	person
	course  string
	college string
}

func main() {
	fmt.Println("Inheritance")

	people := person{"Eduardo", "Dias", 25, 175}

	peopleStudent := student{people, "Analise de sistemas", "USP"}

	fmt.Println(people)
	fmt.Println(peopleStudent)

	fmt.Println(peopleStudent.name)
	fmt.Println(peopleStudent.height)
}
