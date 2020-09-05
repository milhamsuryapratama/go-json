package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
)

var path = "/project/golang/json-data/data.json"

type Todo struct {
	Name string `json:"Name"`
}

func isError(err error) bool {
	if err != nil {
		fmt.Println(err.Error())
	}
	return (err != nil)
}

func createFile() {
	var _, err = os.Stat(path)

	if os.IsNotExist(err) {
		var file, err = os.Create(path)
		if isError(err) {
			return
		}
		defer file.Close()
	}
}

func writeFile() {
	file, err := ioutil.ReadFile(path)
	if err != nil {
		fmt.Println("File tidak ditemukan")
	}

	var input string
	fmt.Println("Input todo : ")
	fmt.Scanf("%s", &input)

	data := []Todo{}

	json.Unmarshal(file, &data)

	newData := &Todo{
		Name: input,
	}

	data = append(data, *newData)

	file, _ = json.MarshalIndent(data, "", " ")
	_ = ioutil.WriteFile(path, file, 0644)

	fmt.Printf("Input %s successfully", input)
}

func readFile() {
	content, err := ioutil.ReadFile(path)
	if err != nil {
		fmt.Println(err.Error())
	}

	var todos []Todo

	err2 := json.Unmarshal(content, &todos)
	if err2 != nil {
		fmt.Println(err2.Error())
	}

	for _, x := range todos {
		fmt.Printf("%s \n", x.Name)
	}
}

func main() {
	createFile()
	writeFile()
	// readFile()
}
