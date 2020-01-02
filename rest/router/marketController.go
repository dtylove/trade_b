package router

import (
	"github.com/gin-gonic/gin"
	"trade_b/matching"
	"trade_b/rest/response"
	"trade_b/utils"
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
