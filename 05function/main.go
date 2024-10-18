package main

import "fmt"

func greeting() {
	fmt.Println("Asalam e alaikummmmm from golang..")
}

func add(a, b int) (result int) {
	result = a + b
	return
}

func mul(a, b int) (result int) {
	result = a * b
	return
}
func main() {
	greeting()
	fmt.Println("here we are learning about fuction...")

	result := add(4, 1)
	fmt.Println("result of add function is :", result)

	result1 := mul(2, 5)
	fmt.Println("result of mul function is :", result1)

}
