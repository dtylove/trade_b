package response

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func Res(ctx *gin.Context, obj interface{}){

	ctx.JSON(http.StatusOK, "参数不正确")
}
