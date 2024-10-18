package main

import "fmt"

// divide1 performs division and returns the result and an error message as a string.
func divide1(a, b float64) (float64, string) {
	if b == 0 {
		return 0, "division by zero is not allowed"
	}
	return a / b, "nil" // Return an empty string for no error
}

// divide function to handle division operation and return error if divisor is zero
//here we used error in the function signature to return error.and we also use string in signature.

// func divide(a, b float64) (float64, error) {
// 	if b == 0 {
// 		return 0, fmt.Errorf("division by zero is not allowed")
// 	}
// 	return a / b, nil
// }

func main() {
	fmt.Println("here we are leaning about how to handle errors in functions.")

	ans, err := divide1(10, 2) // Change the second argument to test different cases
	if err != "nil" {
		fmt.Println("Error occurred while dividing:", err)
		return
	}
	fmt.Println("The division of two numbers is:", ans)

	// 	result, err := divide(10, 0)
	// 	if err != nil {
	// 		fmt.Println("Error occurred while dividing: ", err)
	// 	}
	// 	fmt.Println("The division of two numbers is :", result)
}
