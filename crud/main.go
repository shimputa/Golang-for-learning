package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type Todo struct {
	UserId    int    `json:"userId"`
	Id        int    `json:"id"`
	Title     string `json:"title"`
	Completed bool   `json:"completed"`
}

func main() {
	// here we are learing about crud operations
	res, err := http.Get("https://jsonplaceholder.typicode.com/todos/1")
	if err != nil {
		fmt.Println("error while getting data from API: ", err)
		return
	}
	defer res.Body.Close()

	if res.StatusCode != http.StatusOK {
		fmt.Println("Error while receiving data from API: ", res.StatusCode)
		return
	}

	// data, errerr := io.ReadAll(res.Body)
	// if errerr != nil {
	// 	fmt.Println("error while reading response body: ", errerr)
	// 	return
	// }

	// fmt.Println("Data received: ", string(data))

	fmt.Println("------------------------------------------------------")

	var todo Todo
	err = json.NewDecoder(res.Body).Decode(&todo)
	if err != nil {
		fmt.Println("error while decoding JSON: ", err)
		return
	}
	fmt.Println("Todo received: ", todo)

}
