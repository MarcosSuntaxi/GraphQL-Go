package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/graphql-go/graphql"
	"github.com/graphql-go/handler"
)

// Define the data model
type Task struct {
	ID    string `json:"id"`
	Title string `json:"title"`
	Done  bool   `json:"done"`
}

// Create the Task object type for GraphQL
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

// Sample data
var tasksData = []Task{
	{ID: "1", Title: "Learn GraphQL", Done: false},
	{ID: "2", Title: "Build a GraphQL API in Go", Done: false},
}

// Create the GraphQL query
var rootQuery = graphql.Fields{
	"tasks": &graphql.Field{
		Type: graphql.NewList(taskType),
		Resolve: func(p graphql.ResolveParams) (interface{}, error) {
			return tasksData, nil
		},
	},
}

func main() {
	// Create the GraphQL schema
	schema, err := graphql.NewSchema(graphql.SchemaConfig{
		Query: graphql.NewObject(graphql.ObjectConfig{
			Name:   "RootQuery",
			Fields: rootQuery,
		}),
	})
	if err != nil {
		log.Fatalf("Failed to create schema, %v", err)
	}

	// Set up the GraphQL server with the request handler
	http.Handle("/graphql", handler.New(&handler.Config{
		Schema: &schema,
		Pretty: true,
	}))

	// Start the server on port 8080
	fmt.Println("Server started at http://localhost:8080/graphql")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
