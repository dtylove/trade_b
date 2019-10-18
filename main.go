package main

import (
	"dtyTrade/matching"
	"github.com/gin-gonic/gin"

	//"github.com/shopspring/decimal"
	"time"
)

func main() {
	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	engine := matching.InitEngine(1)
	//go engine.RunOrderFetcher()
	go engine.RunOrderApplier()

	time.Sleep(time.Second * 10)

	r.Run() // listen and serve on 0.0.0.0:8080
}
