package todo

import (
	"net/http"

	"github.com/ant0ine/go-json-rest/rest"
)

func (db *Database) GetTodos(w rest.ResponseWriter, r *rest.Request) {
	todos := Todos{}
	db.DB.Find(&todos)
	w.WriteJson(&todos)
}

func (db *Database) GetTodo(w rest.ResponseWriter, r *rest.Request) {
	id := r.PathParam("id")
	todo := Todo{}
	if db.DB.First(&todo, id).Error != nil {
		rest.NotFound(w, r)
		return
	}
	w.WriteJson(&todo)
}

func (db *Database) PostTodo(w rest.ResponseWriter, r *rest.Request) {
	todo := Todo{}

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

func (db *Database) PutTodo(w rest.ResponseWriter, r *rest.Request) {
	id := r.PathParam("id")
	todo := Todo{}
	if db.DB.First(&todo, id).Error != nil {
		rest.NotFound(w, r)
		return
	}

	updated := Todo{}
	if err := r.DecodeJsonPayload(&updated); err != nil {
		rest.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	todo.Completed = updated.Completed

	if err := db.DB.Save(&todo).Error; err != nil {
		rest.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.WriteJson(&todo)
}

func (db *Database) DeleteTodo(w rest.ResponseWriter, r *rest.Request) {
	id := r.PathParam("id")
	todo := Todo{}
	if db.DB.First(&todo, id).Error != nil {
		rest.NotFound(w, r)
		return
	}

	if err := db.DB.Delete(&todo).Error; err != nil {
		rest.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
}
