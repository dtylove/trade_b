package main

import (
	_ "trade_b/config"
	"trade_b/matching"
	"trade_b/rest"
	"trade_b/rest/models"
)

func main() {

	models.InitDB()

	matching.InitEngineFactory()

	rest.Start()


}
