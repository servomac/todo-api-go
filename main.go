package main

import (
	"log"
	"net/http"
	"time"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"

	"github.com/ant0ine/go-json-rest/rest"
)

func main() {
	db := database{}
	db.InitDB()
	db.InitSchema()

	api := rest.NewApi()
	api.Use(rest.DefaultDevStack...)

	router, err := rest.MakeRouter(
		rest.Get("/todos", db.GetTodos),
		rest.Get("/todos/:id", db.GetTodo),
		rest.Post("/todos", db.PostTodo),
	)
	if err != nil {
		log.Fatal(err)
	}
	api.SetApp(router)
	log.Fatal(http.ListenAndServe(":8080", api.MakeHandler()))
}

type database struct {
	DB *gorm.DB
}

func (db *database) InitDB() {
	var err error
	db.DB, err = gorm.Open("sqlite3", "/tmp/gorm.db")
	if err != nil {
		log.Fatalf("Got an error connecting to database: '%v'", err)
	}
	db.DB.LogMode(true)
}

func (db *database) InitSchema() {
	db.DB.AutoMigrate(&Todo{})
}

type Todo struct {
	ID        int64     `json:"id"`
	Name      string    `json:"name"`
	Completed bool      `json:"completed"`
	Due       time.Time `json:"due"`
}

type Todos []Todo

func (db *database) GetTodos(w rest.ResponseWriter, r *rest.Request) {
	todos := Todos{}
	db.DB.Find(&todos)
	w.WriteJson(&todos)
}

func (db *database) GetTodo(w rest.ResponseWriter, r *rest.Request) {
	id := r.PathParam("id")
	todo := Todo{}
	if db.DB.First(&todo, id).Error != nil {
		rest.NotFound(w, r)
		return
	}
	w.WriteJson(&todo)
}

func (db *database) PostTodo(w rest.ResponseWriter, r *rest.Request) {
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
