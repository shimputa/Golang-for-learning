package main

import "fmt"

func main() {
	fmt.Println("here today we learning about switch statement:")

	day := "sunday"
	switch day {
	case "Monday":
		fmt.Println("It's Monday!")
	case "Tuesday":
		fmt.Println("It's Tuesday!")
	case "Wednesday":
		fmt.Println("It's Wednesday!")
	case "Thursday":
		fmt.Println("It's Thursday!")
	case "Friday":
		fmt.Println("It's Friday!")
	default:
		fmt.Println("It's the weekend!")
	}

	number := 15
	switch {
	case number < 0:
		fmt.Println("The number is negative .")

	case number >= 0 && number <= 10:
		fmt.Println("The number is between 0 and 10 .")

	case number >= 10 && number <= 20:
		fmt.Println("The number is between 11 and 20 .")
	default:
		fmt.Println("The number is greater then 20 .")
	}

	fruit := "apple"
	switch fruit {
	case "apple", "banana", "cherry":
		fmt.Println("This fruit is a common fruit.")
	case "kiwi", "dragonfruit":
		fmt.Println("This fruit is a tropical fruit.")
	default:
		fmt.Println("Unknown fruit type.")
	}

}
