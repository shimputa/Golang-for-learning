package main

import "fmt"

// define a struct
type student struct {
	name  string
	age   int
	grade string
}

func main() {
	// here we will learn about struct in go
	// struct is a collection of fields
	// struct is a collection of data types
	// struct is a collection of values
	// struct is a collection of variables

	//create new student without using new keyword
	student1 := student{"John", 20, "A"}

	fmt.Println("student1: ", student1)

	// create student with var keyword
	var student3 student
	student3.name = "Jane"
	student3.age = 21
	student3.grade = "B"
	fmt.Println("student3: ", student3)

	//create new student using new keyword
	student2 := new(student)
	student2.name = "kane"
	student2.age = 31
	student2.grade = "f"
	fmt.Println("student2: ", student2)
}
