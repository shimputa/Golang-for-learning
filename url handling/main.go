package main

import (
	"fmt"
	"net/url"
)

func main() {
	// learning url handling
	// we will use the net/url package in Go to parse and manipulate URLs

	fmt.Println("Here today we will learn about URL handling")

	myurl := "https://www.example.com/path/to/resource?query=param1&param2=value2"

	parsedUrl, err := url.Parse(myurl)
	if err != nil {
		fmt.Println("Error parsing URL:", err)
		return
	}

	fmt.Println("Protocol:", parsedUrl.Scheme)
	fmt.Println("Host:", parsedUrl.Host)
	fmt.Println("Path:", parsedUrl.Path)
	fmt.Println("rawQuery:", parsedUrl.RawQuery)

	//modify the url
	parsedUrl.Path = "/new/path"
	parsedUrl.RawQuery = "param1=newvalue&param3=value3"
	//constructing the url string from the url object
	newurl := parsedUrl.String()
	fmt.Println("New URL:", newurl)
}
