package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	muxRouter := mux.NewRouter()

	muxRouter.HandleFunc("/users", getUsers).Methods("GET")
	
	fmt.Println("API is on :8000")
	log.Fatal(http.ListenAndServe(":8000", muxRouter))
}

type User struct {
	Id   int    `json:"id"`
	Name string `json:"name"`
}

func getUsers(w http.ResponseWriter, request *http.Request) {
	w.Header().Set("Content=Type", "application/json")
	json.NewEncoder(w).Encode([]User{
		{
			Id:   1,
			Name: "Kirk",
		},
		{
			Id:   2,
			Name: "Spock",
		},
	})
}
