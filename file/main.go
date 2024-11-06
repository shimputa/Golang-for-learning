package main

import (
	"fmt"
	"io"
	"os"
)

func main() {
	file, err := os.Create("example.txt")
	if err != nil {
		fmt.Println("error while creating example.txt: ", err)
		return
	}
	defer file.Close()
	content := "hello world by waseem"
	_, errors := io.WriteString(file, content+"\n")
	if errors != nil {
		fmt.Println("error while writing to file")
		return
	}
	fmt.Println("successfully created example.txt")

}
