package main

import (
	"cql/graph"
	"cql/graph/generated"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/joho/godotenv"
	"github.com/sirupsen/logrus"
)

const defaultPort = "8080"

func loadEnv() (e error) {
	if err := godotenv.Load(); err != nil {
		e = fmt.Errorf("Error loading .env file: %v", err)
	}
	return
}

func init() {
	logrus.SetLevel(logrus.DebugLevel)
}

func main() {
	loadEnv()

	port := os.Getenv("PORT")
	if port == "" {
		port = defaultPort
	}

	srv := handler.NewDefaultServer(generated.NewExecutableSchema(generated.Config{Resolvers: &graph.Resolver{}}))

	http.Handle("/", playground.Handler("GraphQL playground", "/query"))
	http.Handle("/query", srv)

	log.Printf("connect to http://localhost:%s/ for GraphQL playground", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}
