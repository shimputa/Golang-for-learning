package main

import "fmt"

func main() {

	fmt.Println("Hello, today we will learn about map")

	// here intialization of map
	studentGrades := make(map[string]int)
	// Adding key-value pairs to the map
	studentGrades["Alice"] = 92
	studentGrades["Bob"] = 85
	studentGrades["Charlie"] = 78

	// Accessing values from the map
	fmt.Println("Alice's marks:", studentGrades["Alice"])
	fmt.Println("Bob's marks:", studentGrades["Bob"])
	fmt.Println("Charlie's marks:", studentGrades["Charlie"])

	// Modifying values in the map
	studentGrades["Alice"] = 95
	fmt.Println("Alice's marks after modification:", studentGrades["Alice"])

	// Deleting a key-value pair from the map
	delete(studentGrades, "Bob")
	fmt.Println("Bob's marks after deletion:", studentGrades["Bob"])

	// Checking if a key exists in the map
	// here we are checking if the key "Alice" exists in the map
	// if it exists then we are printing the value of the key "Alice"
	// if it does not exist then we are printing that the key "Alice" does not exist
	value, exists := studentGrades["Alice"]
	fmt.Printf("Alice's marks is: %d and  it is exists: %t\n ", value, exists)

	// checking if a key exists in the map using if else
	// here we are checking if the key "Bob" exists in the map
	// if it exists then we are printing the value of the key "Bob"
	// if it does not exist then we are printing that the key "Bob" does not exist
	value, exists = studentGrades["Bob"]
	if exists {
		fmt.Println("Bob's marks:", value)
	} else {
		fmt.Println("Bob's marks do not exist")
	}

	// iterating over the map
	// here we are iterating over the map and printing the key and value of the map
	// here key is the name of the student and value is the grade of the student
	// here we are using range to iterate over the map
	// range is used to iterate over the map, slice, array, string, channel, etc.

	for name, score := range studentGrades {
		fmt.Printf("Grade of %s, is : %d\n", name, score)
	}

}
