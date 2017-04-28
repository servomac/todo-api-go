package main

import (
	"encoding/json"
	"log"
	"net/http"
	"time"

	"github.com/gorilla/mux"
)

type Todo struct {
	Name      string    `json:"name"`
	Completed bool      `json:"completed"`
	Due       time.Time `json:"due"`
}

type Todos []Todo

func GetTodosEndpoint(w http.ResponseWriter, r *http.Request) {
	todos := Todos{
		Todo{Name: "Test"},
		Todo{Name: "Another"},
	}

	json.NewEncoder(w).Encode(todos)
}

func main() {
	router := mux.NewRouter()

	router.HandleFunc("/todos", GetTodosEndpoint).Methods("GET")

	log.Fatal(http.ListenAndServe(":8080", router))
}
