package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strings"
)

type Todo struct {
	UserId    int    `json:"userId"`
	Id        int    `json:"id"`
	Title     string `json:"title"`
	Completed bool   `json:"completed"`
}

func PerformGetRequest() {
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

func PerformPosttRequest() {
	todo := Todo{
		UserId:    1,
		Id:        10,
		Title:     "New Task",
		Completed: false,
	}

	// Convert STRUCT to json data
	jsonData, err := json.Marshal(todo)
	if err != nil {
		fmt.Println("error while marshaling json: ", err)
		return
	}

	// Convert json data to string
	jsonString := string(jsonData)

	// convert json string to json reader
	jsonReader := strings.NewReader(jsonString)
	url1 := "https://jsonplaceholder.typicode.com/todos"
	// Perform POST request
	res, error := http.Post(url1, "application/json", jsonReader)
	if error != nil {
		fmt.Println("Error performing POST request", error)
	}
	defer res.Body.Close()
	// data, errerr := io.ReadAll(res.Body)
	// if errerr != nil {
	// 	fmt.Println("error while reading response body: ", errerr)
	// 	return
	// }
	// fmt.Println("response: ", string(data))

	fmt.Println("response status: ", res.Status)
}
func PerformUpdateRequest() {
	todo := Todo{
		UserId:    1,
		Id:        10,
		Title:     "Updated Task",
		Completed: true,
	}

	// Convert STRUCT to json data
	jsonData, err := json.Marshal(todo)
	if err != nil {
		fmt.Println("error while marshaling json: ", err)
		return
	}

	// Convert json data to string
	jsonString := string(jsonData)

	// convert json string to json reader
	jsonReader := strings.NewReader(jsonString)
	url1 := "https://jsonplaceholder.typicode.com/todos/1"

	// Perform PUT request
	req, errrerr := http.NewRequest(http.MethodPut, url1, jsonReader)
	if errrerr != nil {
		fmt.Println("Error performing PUT request", errrerr)
		return
	}

	req.Header.Set("Content-Type", "application/json")

	client := http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("Error performing request", err)
		return
	}
	defer resp.Body.Close()

	data, errerr := io.ReadAll(resp.Body)
	if errerr != nil {
		fmt.Println("error while reading response body: ", errerr)
		return
	}

	fmt.Println("response: ", string(data))
	fmt.Println("response status: ", resp.Status)

}

func PerformDeleteRequest() {
	url1 := "https://jsonplaceholder.typicode.com/todos/1"

	req, errrerr := http.NewRequest(http.MethodDelete, url1, nil)
	if errrerr != nil {
		fmt.Println("Error performing DELETE request", errrerr)
		return
	}

	client := http.Client{}
	res, err := client.Do(req)
	if err != nil {
		fmt.Println("Error performing request", err)
		return
	}
	defer res.Body.Close()

	// data, errerr := io.ReadAll(res.Body)
	// if errerr != nil {
	// 	fmt.Println("error while reading response body: ", errerr)
	// 	return
	// }
	// fmt.Println("response: ", string(data))
	fmt.Println("response status: ", res.Status)
}

func main() {
	// here we are learing about crud operations
	fmt.Println("learing about crud operations")
	//PerformGetRequest()
	// PerformPosttRequest()
	// PerformUpdateRequest()
	PerformDeleteRequest()

}
