package main

import (
	"log"
	"net/http"
	"time"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"

	"github.com/ant0ine/go-json-rest/rest"
)

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
	db := database{}
	db.InitDB()

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
