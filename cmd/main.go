package main

import (
	"github.com/FaridehGhani/prancer_test/api"
	"github.com/FaridehGhani/prancer_test/deliverpoint/delivery"

	"log"
)

func main() {
	// load agents
	api.Agents = delivery.InitiateAgents(8)

	// run api server
	router := api.Router()
	err := router.Run()
	if err != nil {
		log.Fatalf("http server running error: %v", err)
	}
}
