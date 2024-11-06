package main

import (
	"encoding/json"
	"fmt"
)

//learning about json and how to deal with them

type Person struct {
	Name    string   `json:"name"`
	Age     int      `json:"age"`
	City    string   `json:"city"`
	Hobbies []string `json:"hobbies"`
}

func main() {
	fmt.Println("learning about dealing with json in go")
	GeneratesPersonalData()

}

func GeneratesPersonalData() {
	persons := []Person{
		{Name: "John Doe", Age: 25, City: "New York", Hobbies: []string{"Reading", "Running"}},
		{Name: "Jane Smith", Age: 30, City: "Los Angeles", Hobbies: []string{"Gardening", "Cooking"}},
		{Name: "Michael Johnson", Age: 28, City: "Chicago", Hobbies: []string{"Painting", "Swimming"}},
	}
	// Convert the slice of persons to json format
	jsonData, err := json.MarshalIndent(persons, " ", "  ")
	if err != nil {
		fmt.Println("Error marshalling json:", err)
	}

	fmt.Println("Json data:", string(jsonData))

	//unmarshal
	var unmarshalledPersons []Person
	err = json.Unmarshal(jsonData, &unmarshalledPersons)
	if err != nil {
		fmt.Println("Error unmarshalling json:", err)
	}

	fmt.Println("Unmarshaled data:")
	for _, person := range unmarshalledPersons {
		fmt.Printf("Name: %s, Age: %d, City: %s, Hobbies: %v\n", person.Name, person.Age, person.City, person.Hobbies)
	}
}
