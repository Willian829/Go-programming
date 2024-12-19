package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/users", getUsers)
	fmt.Println("API is on :8000")

	log.Fatal(http.ListenAndServe(":8000", nil))
}

type User struct {
	Id int `json:"id"`
	Name string `json:"name"`
}

func getUsers(w http.ResponseWriter, request *http.Request) {
	if request.Method != "GET" {
		http.Error(w, http.StatusText(http.StatusMethodNotAllowed), http.StatusMethodNotAllowed)
		return
	}
	
	w.Header().Set("Content=Type", "application/json")
	json.NewEncoder(w).Encode([]User{
		{
			Id: 1,
			Name: "Kirk",
		},
		{
			Id: 2,
			Name: "Spock",
		},
	})
} 