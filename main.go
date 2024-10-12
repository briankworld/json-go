package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
)

type ToDo struct {
	ID        int    `json:"id"`
	UserID    int    `json:"userId"`
	Title     string `json:"title"`
	Completed bool   `json:"completed"`
}

func Fetch(url string) (*ToDo, error) {
	resp, err := http.Get(url)
	if err != nil {
		return &ToDo{}, err
	}

	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return &ToDo{}, err
	}

	var toDo ToDo
	err = json.Unmarshal(body, &toDo)
	if err != nil {
		return &ToDo{}, err
	}

	return &toDo, nil
}

func main() {
	apiUrl := "https://jsonplaceholder.typicode.com/todos/1"

	toDo, err := Fetch(apiUrl)
	if err != nil {
		log.Fatalf("Eerror fetching %v", err)
	}

	fmt.Printf("Fetched\n ID: %d\n Title: %s\n Completed %t\n", toDo.ID, (*toDo).Title, toDo.Completed)
}
