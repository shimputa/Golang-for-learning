package main

import "fmt"

func main() {
	//here we are learning about slice
	//slice is a data structure that is used to store a collection of elements of the same type
	//slice is a reference type
	//slice is a dynamic array
	var fruitList = []string{"apple", "banana", "grapes"}
	fruitList = append(fruitList, "mango", "pineapple")
	fmt.Println("fruit list is :", fruitList)
	fmt.Println("length of fruit list is :", len(fruitList))

	// here we are getting the data from index 1 to end
	// fruitList = fruitList[1:]
	// fmt.Println("fruit list is :", fruitList)

	// here we are getting the data from index 1 to 3
	//but not including 3 because of 3 is the length of
	//the slice so it will not include 3 that's why it will print only 2 elements
	// fruitList = fruitList[1:3]
	// fmt.Println("fruit list is :", fruitList)

	// now i want to get the data from slice but not include index 2
	// so i will do this
	// fruitList = append(fruitList[:2], fruitList[3:]...)
	// fmt.Println("fruit list is :", fruitList)

	// here we are creating a slice of length 4 and capacity 8
	highScores := make([]int, 4)
	highScores[0] = 234
	highScores[1] = 235
	highScores[2] = 236
	highScores[3] = 237
	// highScores[4] = 238
	// here we are getting an error because we have only 4 elements in the slice
	// so we need to use append to add more elements to the slice
	highScores = append(highScores, 555, 666, 777, 44, 55, 77, 88, 99)
	fmt.Println("highScores list is :", highScores)

	// here we are removing the element from the slice
	highScores = append(highScores[:2], highScores[3:]...)
	fmt.Println("highScores list is :", highScores)

}
