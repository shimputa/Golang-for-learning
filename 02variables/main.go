package main

import "fmt"

const LoginToken string = "loginheruu"

func main() {
	//variable declared with the var keyword
	var name string = "waseem"
	var age uint8 = 25

	fmt.Println("name: ", name)
	fmt.Printf("name: %T \n", name)
	fmt.Println("age: ", age)
	fmt.Printf("age: %T \n", age)

	//constant values
	fmt.Println("LoginToken: ", LoginToken)
	fmt.Printf("LoginToken: %T \n", LoginToken)

	// variable declared with the var keyword and initialized with a default value
	var isLoggedIn bool = false
	fmt.Println("isLoggedIn: ", isLoggedIn)
	fmt.Printf("isLoggedIn: %T \n", isLoggedIn)

	// variable declared with the var keyword and initialized with a default value
	var isEnabled bool = true
	fmt.Println("isEnabled: ", isEnabled)
	fmt.Printf("isEnabled: %T \n", isEnabled)

	// variable declared with the var keyword and initialized with a default value
	var isReady bool = true
	fmt.Println("isReady: ", isReady)
	fmt.Printf("isReady: %T \n", isReady)

	// variable declared

	// variable declared with the := operator
	city := "karachi"
	fmt.Println("city: ", city)
	fmt.Printf("city: %T \n", city)

	// variable declared with the var keyword and initialized later
	var greeting string
	greeting = "Hello, World!"
	fmt.Println("greeting: ", greeting)
	fmt.Printf("greeting: %T \n", greeting)
}
