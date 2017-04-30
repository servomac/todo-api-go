package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/ant0ine/go-json-rest/rest"

	"github.com/servomac/todo-api-go/bundles/todo"
)

func main() {
	config := getConfig()

	db := todo.Database{}
	db.InitDB(config.Database)
	db.InitSchema()

	api := rest.NewApi()
	api.Use(rest.DefaultDevStack...)
	api.Use(&rest.CorsMiddleware{
		RejectNonCorsRequests: false,
		OriginValidator: func(origin string, request *rest.Request) bool {
			return origin == "http://localhost:3000"
		},
		AllowedMethods: []string{"GET", "POST", "PUT"},
		AllowedHeaders: []string{
			"Accept", "Content-Type", "X-Custom-Header", "Origin"},
		AccessControlAllowCredentials: true,
		AccessControlMaxAge:           3600,
	})

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

	log.Fatal(initServer(config, api))
}

func initServer(c Config, api *rest.Api) error {
	addr := fmt.Sprintf(":%v", c.Port)
	return http.ListenAndServe(addr, api.MakeHandler())
}
