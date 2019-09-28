package order

import "github.com/shopspring/decimal"

type Order struct {
	Id         int64           // 订单id
	IsBuy      bool            // 买true bids 卖false asks
	BookId     int             // order book id (针对不同市场 货币对)
	OrderType  string          // asks bids
	Price      decimal.Decimal // 价格
	Volume     decimal.Decimal // 总价值 总个数*价格
	Quantity   decimal.Decimal // 总个数
	Remained   decimal.Decimal // 剩余总量
	MatchVol   decimal.Decimal // 本次交易量
	MatchCount int             // 成交次数
	Timestamp  int64           // 创建时间戳
	UserId     int64           // 用户id
	IsRemain   bool            // 是否完全成交
}

//type Trade struct {
//	MatchVol decimal.Decimal // 本次交易量
//	AskId    int64           // 卖家id
//	BidsId   int64           //卖家id
//}

// 是否可以成交
func (order *Order) IsCrossed(price decimal.Decimal) (result bool) {
	if order.IsBuy == false {
		result = price.GreaterThan(order.Price)
	} else {
		result = price.LessThan(order.Price)
	}
	return
}

// 交易
func (order *Order) TradeWith(counterOrder Order) {

	matchVol := decimal.Min(order.Volume, counterOrder.Volume)
	order.Volume = order.Volume.Sub(matchVol)
	counterOrder.Volume = counterOrder.Volume.Sub(matchVol)

	return
}
