package main

import (
	_ "dtyTrade/config"
	"dtyTrade/matching"
	"dtyTrade/rest"
	"dtyTrade/rest/models"
)

func main() {

	models.InitDB()

	matching.InitEngineFactory()

	rest.Start()


}
