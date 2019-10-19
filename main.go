package main

import (
	"dtyTrade/matching"
	"dtyTrade/router"
)

func main() {

	engine := matching.InitEngine(1)
	//go engine.RunOrderFetcher()
	go engine.RunOrderApplier()

	router.Start()
}
