package main

import (
	"log"
	"net/http"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/rlarkin212/bjj-cs/cmd/gql/graph"
	"github.com/rlarkin212/bjj-cs/cmd/gql/graph/generated"
	"github.com/rlarkin212/bjj-cs/configs"
	"github.com/rlarkin212/bjj-cs/internal/service/search"
	"github.com/rlarkin212/bjj-cs/internal/service/submit"
)

const defaultPort = "8080"

func main() {
	config, err := configs.LoadConfig("./../../configs/", "config", "yaml")
	if err != nil {
		log.Fatal(err)
	}

	port := config.GQL.Port
	if port == "" {
		port = defaultPort
	}

	resolver := &graph.Resolver{
		SearchService: search.NewSearchService(config),
		SubmitService: submit.NewSubmitService(config),
	}

	srv := handler.NewDefaultServer(generated.NewExecutableSchema(generated.Config{Resolvers: resolver}))

	http.Handle("/", playground.Handler("GraphQL playground", "/query"))
	http.Handle("/query", srv)

	log.Printf("connect to http://localhost:%s/ for GraphQL playground", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}
