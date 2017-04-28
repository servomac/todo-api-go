package main

import (
	"log"
	"net/http"

	"github.com/ant0ine/go-json-rest/rest"

	"github.com/servomac/goapi/bundles/todo"
)

func main() {
	db := database{}
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

func (db *database) GetTodos(w rest.ResponseWriter, r *rest.Request) {
	todos := todo.Todos{}
	db.DB.Find(&todos)
	w.WriteJson(&todos)
}

func (db *database) GetTodo(w rest.ResponseWriter, r *rest.Request) {
	id := r.PathParam("id")
	todo := todo.Todo{}
	if db.DB.First(&todo, id).Error != nil {
		rest.NotFound(w, r)
		return
	}
	w.WriteJson(&todo)
}

func (db *database) PostTodo(w rest.ResponseWriter, r *rest.Request) {
	todo := todo.Todo{}

	if err := r.DecodeJsonPayload(&todo); err != nil {
		rest.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	if err := db.DB.Save(&todo).Error; err != nil {
		rest.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(201)
	w.WriteJson(&todo)
}

func (db *database) PutTodo(w rest.ResponseWriter, r *rest.Request) {
	id := r.PathParam("id")
	original := todo.Todo{}
	if db.DB.First(&original, id).Error != nil {
		rest.NotFound(w, r)
		return
	}

	updated := todo.Todo{}
	if err := r.DecodeJsonPayload(&updated); err != nil {
		rest.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	original.Completed = updated.Completed

	if err := db.DB.Save(&original).Error; err != nil {
		rest.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.WriteJson(&original)
}

func (db *database) DeleteTodo(w rest.ResponseWriter, r *rest.Request) {
	id := r.PathParam("id")
	original := todo.Todo{}
	if db.DB.First(&original, id).Error != nil {
		rest.NotFound(w, r)
		return
	}

	if err := db.DB.Delete(&original).Error; err != nil {
		rest.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
}
