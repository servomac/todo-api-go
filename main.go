package main

import (
	"log"
	"net/http"

	"github.com/ant0ine/go-json-rest/rest"

	"github.com/servomac/goapi/bundles/todo"
)

func main() {
	db := todo.Database{}
	db.InitDB()
	db.InitSchema()

	api := rest.NewApi()
	api.Use(rest.DefaultDevStack...)

	router, err := rest.MakeRouter(
		rest.Get("/todos", db.GetTodos),
		rest.Post("/todos", db.PostTodo),
		rest.Get("/todos/:id", db.GetTodo),
		rest.Put("/todos/:id", db.PutTodo),
		rest.Delete("/todos/:id", db.DeleteTodo),
	)
	if err != nil {
		log.Fatal(err)
	}
	api.SetApp(router)
	log.Fatal(http.ListenAndServe(":8080", api.MakeHandler()))
}
