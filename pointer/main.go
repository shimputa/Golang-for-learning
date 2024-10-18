package main

import (
	"fmt"
)

func updateNum(num *int) {
	*num = *num + 1
}

func swap(a, b *int) {
	temp := *a
	*a = *b
	*b = temp
}

func main() {

	num := 5
	fmt.Println("The num Before updated ", num)
	updateNum(&num)
	fmt.Println("The num After updated ", num)

	a, b := 10, 20
	fmt.Printf("Before Swaping: a = %d, b = %d\n", a, b)
	swap(&a, &b)
	fmt.Printf("After Swaping: a =%d, b=%d\n", a, b)

}
