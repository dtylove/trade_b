package models

import (
	"time"
)

type Order struct {
	Id        uint `gorm:"primary_key"` // 订单id
	CreatedAt time.Time
	UpdatedAt time.Time

	Price    string // 价格
	Quantity string // 总个数
	Remained string // 剩余个数
	MatchQ   string // 本次交个数

	IsBuy      bool                       // 买true bids 卖false asks
	OrderType  string                     // asks bids
	MarketId   uint                       // order book id (针对不同市场 货币对)
	MatchCount uint                       // 成交次数
	Timestamp  uint                       // 创建时间戳
	UserId     uint                       // 用户id
	IsRemain   bool `gorm:"default:true"` // 是否完全成交
}
