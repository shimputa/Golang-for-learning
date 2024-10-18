package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	// Declare and initialize a string variable with a welcome message
	welcome := "Welcome to user input"
	// Print the welcome message to the console
	fmt.Println(welcome)
	// Create a new reader for standard input
	reader := bufio.NewReader(os.Stdin)
	// Prompt the user to enter their name and read the input line by line, excluding the newline character
	fmt.Print("Enter your name: ")
	// Read a line from standard input, excluding the newline character
	input, _ := reader.ReadString('\n')
	// Trim the newline character from the end of the input
	fmt.Printf("Hello, %s", input)
	// Print a confirmation message to the console
	fmt.Println("thanks you :", input)
	// Print the type of the name variable to verify its data type
	fmt.Printf("the type of name is %T \n", input)

}
