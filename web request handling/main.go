package main

import (
	"fmt"
	"io"
	"net/http"
)

func main() {
	// learning about web request handling
	// we will use the net/http package in Go to create a simple web server

	fmt.Println("Here today we will learn about web request handling")

	res, err := http.Get("https://jsonplaceholder.typicode.com/todos/1")
	if err != nil {
		fmt.Println("Error getting Get response", err)
	}

	defer res.Body.Close() // Close the response body when we're done

	data, err := io.ReadAll(res.Body)
	if err != nil {
		fmt.Println("Error reading response body", err)
	}

	fmt.Println("Response Body:", string(data)) // Print the response body as a string

}
