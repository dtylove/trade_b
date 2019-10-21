package main

import (
	_ "dtyTrade/config"
	"dtyTrade/matching"
	"dtyTrade/models"
	"dtyTrade/router"
)

func main() {

	engine := matching.InitEngine(1)
	//go engine.RunOrderFetcher()
	go engine.RunOrderApplier()

	models.InitDB()
	router.Start()
}
