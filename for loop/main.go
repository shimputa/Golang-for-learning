package main

import "fmt"

func main() {
	fmt.Println("Hello, World! today we will learn about for loop")

	// for loop with initialization, condition, and increment
	for i := 0; i < 5; i++ {
		fmt.Println("value of i is", i)
	}

	// for loop with infinite loop
	// infinite loop is a loop that runs indefinitely until a specific condition is met
	count := 0
	for {
		fmt.Println("infinite loop")
		count++
		if count == 5 {
			break
		}
	}

	// for loop with range
	// range is used to iterate over a collection of elements
	numbers := []int{11, 12, 13, 14, 15}
	for index, value := range numbers {
		fmt.Printf("On index %d the value is %d\n", index, value)
	}

	//for loop with range on string
	message := "Here we are learning about for loop"
	for index, value := range message {
		fmt.Printf("On index %d the character is %c\n", index, value)
	}
}
