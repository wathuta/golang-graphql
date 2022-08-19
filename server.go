package main

import (
	"log"
	"net/http"
	"os"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/wathuta/golang-graphql/graph"
	"github.com/wathuta/golang-graphql/graph/generated"
	"github.com/wathuta/golang-graphql/internal/pkg/db"
)

const defaultPort = "8080"

func init(){
	db.Connect()
}
func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = defaultPort
	}
	log.Println(port)
	defer db.CloseDb()
	srv := handler.NewDefaultServer(generated.NewExecutableSchema(generated.Config{Resolvers: &graph.Resolver{}}))

	http.Handle("/", playground.Handler("GraphQL playground", "/query"))
	http.Handle("/query", srv)

	log.Printf("connect to http://localhost:%s/ for GraphQL playground", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}
