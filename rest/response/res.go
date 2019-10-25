package response

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func Res(ctx *gin.Context, msg string, data interface{}) {
	ctx.JSON(http.StatusOK, BuildMsg(msg, data))
}
