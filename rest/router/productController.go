package router

import (
	"github.com/gin-gonic/gin"
	"trade_b/matching"
	"trade_b/rest/models"
	"trade_b/rest/response"
	"trade_b/utils"
)

func GetProduct(ctx *gin.Context) {
	id := ctx.Param("id")
	var pId uint
	if err := utils.StrToUint(id, &pId); err != nil {
		response.Res(ctx, response.C_PARAMS_ERR, nil)
		return
	}

	product := models.Product{Id: pId}

	err := models.FindOne(&product)
	if product.Id == 0 || err != nil {
		response.Res(ctx, response.P_NOT_FOUND, nil)
		return
	}

	response.Res(ctx, response.OK, product)
}

type productRequest struct {
	Base      string // 买方货币
	BaseMinQ  string // 单笔最小
	BaseMaxQ  string // 单笔最大
	BaseScale int32  // 最大小数位

	Counter      string // 卖方货币
	CounterMinQ  string // 单笔最小金额
	CounterMaxQ  string // 单笔最大
	CounterScale int32  // 最大小数位
}

func CreateProduct(ctx *gin.Context) {
	var body productRequest
	if err := ctx.BindJSON(&body); err != nil {
		response.Res(ctx, response.C_PARAMS_ERR, nil)
		return
	}

	if body.Base == body.Counter {
		response.Res(ctx, response.P_BASE_COUNTER_SAME, nil)
		return
	}

	product := models.Product{
		Base:      body.Base,
		BaseMinQ:  body.BaseMinQ,
		BaseMaxQ:  body.BaseMaxQ,
		BaseScale: body.BaseScale,

		Counter:      body.Counter,
		CounterMinQ:  body.CounterMinQ,
		CounterMaxQ:  body.CounterMaxQ,
		CounterScale: body.CounterScale,
	}

	if err := models.Add(&product); err != nil || product.Id == 0 {
		response.Res(ctx, response.P_CREATE_FAILED, err)
		return
	}

	matching.AddProduct(&product)
	response.Res(ctx, response.OK, product)
}

func GetProducts(ctx *gin.Context) {
	var pList []models.Product

	err := models.FindList(&pList)
	if err != nil {
		response.Res(ctx, response.P_NOT_FOUND, nil)
		return
	}

	response.Res(ctx, response.OK, pList)
}
