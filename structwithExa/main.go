package main

import "fmt"

type person struct {
	firstName string
	lastName  string
	age       int
}

type contact struct {
	email string
	phone string
}
type address struct {
	street  string
	city    string
	country string
}
type employee struct {
	employee_detail  person
	employee_contact contact
	employee_address address
}

func main() {
	// create employee with var keyword
	var emp1 employee
	emp1.employee_detail = person{"John", "Doe", 30}
	emp1.employee_contact = contact{"john@gmail.com", "1234567890"}
	emp1.employee_address = address{"123 Main St", "New York", "USA"}
	fmt.Println("Employee 1: ", emp1)

	//create employee without using new keyword
	emp2 := employee{
		employee_detail:  person{"John", "Doe", 30},
		employee_contact: contact{"john@gmail.com", "1234567890"},
		employee_address: address{"123 Main St", "New York", "USA"},
	}
	fmt.Println("Employee 1: ", emp2)

	//create employee using new keyword
	emp3 := new(employee)
	emp3.employee_detail = person{"John", "Doe", 30}
	emp3.employee_contact = contact{"john@gmail.com", "1234567890"}
	emp3.employee_address = address{"123 Main St", "New York", "USA"}
	fmt.Println("Employee 3: ", emp3)
}
