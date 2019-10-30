package router

import (
	"dtyTrade/matching"
	"dtyTrade/rest/response"
	"dtyTrade/utils"
	"fmt"
	"github.com/gin-gonic/gin"
)

func GetMarket(ctx *gin.Context) {
	fmt.Println("call GetMarket ")
	id := ctx.Param("id")
	var pId uint
	if err := utils.StrToUint(id, &pId); err != nil {
		response.Res(ctx, response.C_PARAMS_ERR, nil)
		return
	}

	engine := matching.GetEngine(pId)

	engine.Print(10)
}
