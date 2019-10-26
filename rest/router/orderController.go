package router

import (
	"dtyTrade/rest/models"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/shopspring/decimal"
	"time"

	"dtyTrade/rest/response"
)

type OrderRequest struct {
	Price    string // 单价
	Quantity string // 数量
	MarketId uint   // 市场id
	IsBuy    bool   // true 买 false 卖
}

func SubmitOrder(ctx *gin.Context) {
	var body OrderRequest
	if err := ctx.BindJSON(&body); err != nil {
		response.Res(ctx, response.C_PARAMS_ERR, nil)
		return
	}

	user := ctx.MustGet("user").(*models.User)
	price, err := decimal.NewFromString(body.Price)
	if err != nil {
		response.Res(ctx, response.C_PARAMS_ERR, "price is not number")
		return
	}

	quantity, err := decimal.NewFromString(body.Quantity)
	if err != nil {
		response.Res(ctx, response.C_PARAMS_ERR, "quantity is not number")
		return
	}

	order := models.Order{
		IsBuy:     body.IsBuy,
		MarketId:  body.MarketId,
		Price:     price.String(),
		Quantity:  quantity.String(),
		UserId:    user.Id,
		Timestamp: uint(time.Now().Unix()),
	}

	if err := order.Add(); err != nil {
		fmt.Println(err)
		response.Res(ctx, response.O_ADD_ERR, nil)
		return
	}

	response.Res(ctx, response.OK, order)

}
