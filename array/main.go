package main

import "fmt"

func main() {
	fruits := [5]string{"apple", "", "banana", "mango", "orange"}
	arr := [5]int{1, 2, 3, 5}
	fmt.Println("fruits are:", fruits)
	// %q is used to print the values in quotes
	fmt.Printf("fruits are:%q\n", fruits)
	fmt.Println("Array:", arr)

	fmt.Println("Length of Array:", len(arr))
}
