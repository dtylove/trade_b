package router

import (
	"github.com/gin-gonic/gin"
	"io"
	"os"
)

func Start() {
	gin.SetMode(gin.DebugMode)
	// 打印log 到文件和控制台
	f, _ := os.Create("gin.log")
	gin.DefaultWriter = io.MultiWriter(f, os.Stdout)
	router := gin.Default()

	router.POST("/signup", SignUp)

	err := router.Run()
	if err != nil {
		panic(err)
	}
}
