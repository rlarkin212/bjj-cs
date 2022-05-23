package main

import (
	"log"

	"github.com/rlarkin212/bjj-cs/cmd/rest/api"
	"github.com/rlarkin212/bjj-cs/cmd/rest/routes"
	"github.com/rlarkin212/bjj-cs/configs"
)

func main() {
	config, err := configs.LoadConfig("./../../configs/", "config", "yaml")
	if err != nil {
		log.Fatal(err)
	}

	router := api.Generate()
	routes.RegisterRoutes(router, config)

	api.Start(router, config)
}
