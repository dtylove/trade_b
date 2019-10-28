package router

import (
	"dtyTrade/matching"
	"dtyTrade/rest/models"
	"dtyTrade/rest/response"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/shopspring/decimal"
	"time"
)

type orderRequest struct {
	Price    string // 单价
	Quantity string // 数量
	MarketId uint   // 市场id
	IsBuy    bool   // true 买 false 卖
}
func SubmitOrder(ctx *gin.Context) {
	var body orderRequest
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
		response.Res(ctx, response.C_PARAMS_ERR, "quantity is not string number")
		return
	}

	order := models.Order{
		Price:     body.Price,
		Quantity:  body.Quantity,
		Order: matching.Order{
			IsBuy:     body.IsBuy,
			MarketId:  body.MarketId,
			Price:     price,
			Quantity:  quantity,
			UserId:    user.Id,
			Timestamp: uint(time.Now().Unix()),
		},
	}

	if err := order.Add(); err != nil {
		fmt.Println(err)
		response.Res(ctx, response.O_ADD_ERR, nil)
		return
	}

	response.Res(ctx, response.OK, order)

}
