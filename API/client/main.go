package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

type User struct {
	Id int `json:"id"`
	Name string `json:"name"`
}

func main() {
	resp, err := http.Get("http://localhost:8000/users")
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	if resp.StatusCode != 200 {
		fmt.Println("Not success", resp.StatusCode)
		return
	}

	body, err := io.ReadAll(resp.Body)

	var response []User
	err = json.Unmarshal(body, &response)
	if err != nil {
		fmt.Println("Erro retrieving data", err.Error())
	}

	fmt.Println(response)
}
