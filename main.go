package main

import (
	"log"
	"net/http"
	"time"

	"github.com/ant0ine/go-json-rest/rest"
)

type Todo struct {
	Name      string    `json:"name"`
	Completed bool      `json:"completed"`
	Due       time.Time `json:"due"`
}

type Todos []Todo

func GetTodosEndpoint(w rest.ResponseWriter, r *rest.Request) {
	todos := Todos{
		Todo{Name: "Test"},
		Todo{Name: "Another"},
	}

	w.WriteJson(todos)
}

func main() {
	api := rest.NewApi()
	api.Use(rest.DefaultDevStack...)

	router, err := rest.MakeRouter(
		rest.Get("/todos", GetTodosEndpoint),
	)
	if err != nil {
		log.Fatal(err)
	}
	api.SetApp(router)
	log.Fatal(http.ListenAndServe(":8080", api.MakeHandler()))
}
