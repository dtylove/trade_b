package matching

import (
	"encoding/json"
	"fmt"
	"github.com/shopspring/decimal"
)

var Zero = decimal.NewFromFloat(0.0)

type Order struct {
	Id        uint `gorm:"primary_key"` // 订单id
	IsBuy     bool                      // 买true bids 卖false asks
	OrderType string                    // asks bids
	MarketId  uint                      // order book id (针对不同市场 货币对)

	Price    decimal.Decimal // 价格
	Quantity decimal.Decimal // 总个数
	Remained decimal.Decimal // 剩余个数
	MatchQ   decimal.Decimal // 本次交个数

	MatchCount uint // 成交次数
	Timestamp  uint // 创建时间戳
	UserId     uint // 用户id
	IsRemain   bool // 是否完全成交
}

// 是否可以成交
func (order *Order) IsCrossed(price decimal.Decimal) (result bool) {
	if order.IsBuy == false {
		result = price.GreaterThanOrEqual(order.Price)
	} else {
		result = price.LessThanOrEqual(order.Price)
	}
	return
}

// 交易
func (order *Order) TradeWith(marketOrder *Order) (trade Trade) {
	matchQ := decimal.Min(order.Remained, marketOrder.Remained)
	order.MatchQ = matchQ
	order.Remained = order.Remained.Sub(matchQ)
	order.MatchCount++
	if order.Remained.Equal(Zero) {
		order.IsRemain = false
	}

	marketOrder.MatchQ = matchQ
	marketOrder.Remained = marketOrder.Remained.Sub(matchQ)
	marketOrder.MatchCount++
	if marketOrder.Remained.Equal(Zero) {
		marketOrder.IsRemain = false
	}

	trade.Price = marketOrder.Price
	trade.Funds = matchQ.Mul(trade.Price)
	trade.MatchQ = matchQ

	trade.Taker.OrderId = order.Id
	trade.Taker.Pricing = order.Price
	trade.Taker.MatchQ = matchQ
	trade.Taker.Funds = trade.Funds

	trade.Maker.OrderId = marketOrder.Id
	trade.Maker.Pricing = marketOrder.Price
	trade.Maker.MatchQ = matchQ
	trade.Maker.Funds = trade.Funds

	return
}

func (order *Order) Print() {
	oJson, _ := json.Marshal(order)
	fmt.Println(string(oJson))
}
