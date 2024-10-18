package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	// Print the welcome message to the console
	fmt.Println("welcome to our pizza app")

	// Declare and initialize a variable with a value of 42 and print it to the console
	num := 42
	num1 := "42"
	num2 := 42

	//here whats we do is convert the num to float type
	data := float64(num2)
	data = data + 1.2
	fmt.Println("data is ", data)
	fmt.Printf("type of data is %T\n", data)

	// here i have convert the integer to string type.
	num_str := strconv.Itoa(num)
	fmt.Println("num_str is ", num_str)
	fmt.Printf("type of num_str is %T\n", num_str)

	// Here I have converted the string to integer type.
	num_int, err := strconv.Atoi(num1) // Use Atoi to convert string to integer
	if err != nil {
		fmt.Println("Error converting string to integer:", err)
		return // Handle the error appropriately
	}
	fmt.Println("num_int is ", num_int)
	fmt.Printf("type of num_int is %T\n", num_int)

	// Ask the user for rating and store it in a variable
	fmt.Println("Please rate our pizza between 1 and 5")
	// Read the user input from the console and convert it to a float64
	reader := bufio.NewReader(os.Stdin)
	input, _ := reader.ReadString('\n')
	fmt.Println("thanks for rateing:", input)
	// Add 1 to the user's rating and print the new total
	// Trim the newline character from the input and convert it to a float64 using strconv.ParseFloat() function
	// If the conversion fails, print the error message and exit the program
	// Otherwise, add 1 to the user's rating and print the new total.
	// The input is trimmed of leading and trailing whitespace using strings.TrimSpace() function before conversion.
	// The result is then printed to the console.
	// Note: This solution assumes that the user enters a valid rating (a number between 1 and 5).
	// If the user enters an invalid input, the program will print an error message and exit.
	// Note: The program does not handle the case where the user enters a non-numeric input.
	// If the user enters a non-numeric input, the program will print an error message and exit.

	// here we take input from user and convert it to float64 type.
	numRating, err := strconv.ParseFloat(strings.TrimSpace(input), 64)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("Added 1 to your rating :", numRating+1)
	}
}
