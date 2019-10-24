package main

import (
	_ "dtyTrade/config"
	"dtyTrade/matching"
	"dtyTrade/rest"
	"dtyTrade/rest/models"
)

func main() {

	engine := matching.InitEngine(1)
	//go engine.RunOrderFetcher()
	go engine.RunOrderApplier()

	models.InitDB()
	rest.Start()
}
