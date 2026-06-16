package main

import "fmt"

type Person struct {
	FirstName string
	LastName  string
	Age       int
}

func MakePersonPointer(firstName string, lastName string, age int) *Person {
	return &Person{
		FirstName: firstName,
		LastName:  lastName,
		Age:       age,
	}
}

func MakePerson(firstName string, lastName string, age int) Person {
	return Person{
		FirstName: firstName,
		LastName:  lastName,
		Age:       age,
	}
}

func main() {
	p1 := MakePerson("David", "Johns", 22)
	fmt.Println(p1)
	p2 := MakePersonPointer("John", "Doe", 71)
	fmt.Println(p2)
}
