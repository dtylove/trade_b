package models

import (
	"time"
)

type Order struct {
	Id        uint `gorm:"primary_key"` // 订单id
	IsBuy     bool                      // 买true bids 卖false asks
	OrderType string                    // asks bids
	MarketId  uint                      // order book id (针对不同市场 货币对)

	Price    string `gorm:"type:varchar(64);default:null;" json:"price"`    // 价格
	Quantity string `gorm:"type:varchar(64);default:null;" json:"quantity"` // 总个数
	Remained string `gorm:"type:varchar(64);default:null;" json:"remained"` // 剩余个数
	MatchQ   string `gorm:"type:varchar(64);default:null;" json:"match_q"`  // 本次交个数

	MatchCount uint // 成交次数
	Timestamp  uint // 创建时间戳
	UserId     uint // 用户id
	IsRemain   bool // 是否完全成交
	CreatedAt  time.Time
	UpdatedAt  time.Time
}

func (o *Order) Add() error {
	return GetDB().Create(o).Error
}
