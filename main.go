package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/graphql-go/graphql"
	"github.com/graphql-go/handler"
)

// Definir el modelo de datos
type Task struct {
	ID    string `json:"id"`
	Title string `json:"title"`
	Done  bool   `json:"done"`
}

// Crear el tipo de objeto Task para GraphQL
var taskType = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "Task",
		Fields: graphql.Fields{
			"id":    &graphql.Field{Type: graphql.String},
			"title": &graphql.Field{Type: graphql.String},
			"done":  &graphql.Field{Type: graphql.Boolean},
		},
	},
)

// Crear los datos de ejemplo
var tasksData = []Task{
	{ID: "1", Title: "Learn GraphQL", Done: false},
	{ID: "2", Title: "Build a GraphQL API in Go", Done: false},
}

// Crear la consulta de GraphQL
var rootQuery = graphql.Fields{
	"tasks": &graphql.Field{
		Type: graphql.NewList(taskType),
		Resolve: func(p graphql.ResolveParams) (interface{}, error) {
			return tasksData, nil
		},
	},
}

func main() {
	// Crear el esquema GraphQL
	schema, err := graphql.NewSchema(graphql.SchemaConfig{
		Query: graphql.NewObject(graphql.ObjectConfig{
			Name:   "RootQuery",
			Fields: rootQuery,
		}),
	})
	if err != nil {
		log.Fatalf("Failed to create schema, %v", err)
	}

	// Crear el servidor GraphQL con el manejador de solicitudes
	http.Handle("/graphql", handler.New(&handler.Config{
		Schema: &schema,
		Pretty: true,
	}))

	// Iniciar el servidor en el puerto 8080
	fmt.Println("Server started at http://localhost:8080/graphql")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
