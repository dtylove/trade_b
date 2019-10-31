package matching

import "github.com/shopspring/decimal"

const Asks = "asks"
const Bids = "bids"

type Quote struct {
	Price        decimal.Decimal // 单价
	TotalQ       decimal.Decimal // 当前价格总个数
	Accumulation decimal.Decimal // 累加总个数
}
