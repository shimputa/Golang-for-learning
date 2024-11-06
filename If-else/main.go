package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	fmt.Println("Hello, World! today we are learning about if else")

	//if else is a conditional statement that allows us to execute different blocks of code based on certain conditions.
	//if else is used to make decisions in a program.
	//let start with simple if else program
	age := 18
	if age >= 18 {
		fmt.Println("You are eligible to vote")
	} else {
		fmt.Println("You are not eligible to vote")
	}

	// Method 1: Using fmt.Sscanf
	// this method is less commonly used because it is less readable and harder to understand.
	fmt.Println("\n--- Method 1: Using fmt.Sscanf ---")
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Enter your age: ")
	input, _ := reader.ReadString('\n')
	var userAge int
	fmt.Sscanf(input, "%d", &userAge)

	if userAge >= 18 {
		fmt.Println("You are eligible to vote")
	} else {
		fmt.Println("You are not eligible to vote")
	}

	// Method 2: Using strconv.Atoi (more commonly used)
	// this method is more commonly used because it is more readable and easier to understand.
	fmt.Println("\n--- Method 2: Using strconv.Atoi ---")
	reader = bufio.NewReader(os.Stdin)
	fmt.Print("Enter your age again: ")
	input, _ = reader.ReadString('\n')
	// TrimSpace removes whitespace including newline
	userAge2, err := strconv.Atoi(strings.TrimSpace(input))
	if err != nil {
		fmt.Println("Please enter a valid number!")
		return
	}

	if userAge2 >= 18 {
		fmt.Println("You are eligible to vote")
	} else {
		fmt.Println("You are not eligible to vote")
	}
}
