package rest

import (
	_ "dtyTrade/rest/response"

	"dtyTrade/rest/midware"
	"dtyTrade/rest/router"
	"github.com/gin-gonic/gin"
	"io"
	"os"
)

func Start() {
	gin.SetMode(gin.DebugMode)
	// 打印log 到文件和控制台
	f, _ := os.Create("gin.log")
	gin.DefaultWriter = io.MultiWriter(f, os.Stdout)
	r := gin.Default()

	userGroup := r.Group("/user")
	userGroup.POST("/signup", router.SignUp)
	userGroup.POST("/signin", router.SignIn)
	userGroup.GET("/:id", midware.VerifyToken(), router.GetUser)

	orderGroup := r.Group("/order")
	orderGroup.POST("/submit", midware.VerifyToken(), router.SubmitOrder)

	productGroup := r.Group("/product")
	productGroup.GET("/single/:id", router.GetProduct)
	productGroup.GET("/list", router.GetProducts)
	productGroup.POST("/create", router.CreateProduct)

	marketGroup := r.Group("/market")
	marketGroup.GET("/quote/:id", router.GetMarket)

	err := r.Run()
	if err != nil {
		panic(err)
	}
}
