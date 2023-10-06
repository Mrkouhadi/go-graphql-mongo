package main

import (
	"log"
	"net/http"

	"go-gql-mongo/graph"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
)

// FIXME: enable CORS so we can access it using javascript
// chi router + graphql : https://blog.devgenius.io/intergrating-graphql-with-golang-the-right-way-d6a27bf4cbf7
func main() {
	port := "8080"
	var srv = handler.NewDefaultServer(graph.NewExecutableSchema(graph.Config{Resolvers: &graph.Resolver{}}))
	http.Handle("/", playground.Handler("GraphQL playground", "/query"))
	http.Handle("/query", srv)
	log.Printf("connect to http://localhost:%s/ for GraphQL playground", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}

/// enabling CORS:
// func enableCors(w *http.ResponseWriter) {
// 	(*w).Header().Set("Access-Control-Allow-Origin", "*") // this means open all possible origins
// 	// (*w).Header().Set("Access-Control-Allow-Origin", "http://localhost:3000") // open to only http://localhost:3000
// }
