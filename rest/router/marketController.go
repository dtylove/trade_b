package router

import (
	"dtyTrade/matching"
	"dtyTrade/rest/response"
	"dtyTrade/utils"
	"github.com/gin-gonic/gin"
)

type depthResp struct {
	Asks []matching.Quote `json:"asks"`
	Bids []matching.Quote `json:"bids"`
}

func GetMarket(ctx *gin.Context) {
	id := ctx.Param("id")
	var pId uint
	if err := utils.StrToUint(id, &pId); err != nil {
		response.Res(ctx, response.C_PARAMS_ERR, nil)
		return
	}

	engine := matching.GetEngine(pId)
	asks, bids := engine.GetDepth()
	depth := depthResp{
		Asks: asks,
		Bids: bids,
	}

	response.Res(ctx, response.OK, depth)
}
