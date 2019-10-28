package models

import (
	"dtyTrade/matching"
	"time"
)

type Order struct {
	matching.Order
	Price     string // 价格
	Quantity  string // 总个数
	CreatedAt time.Time
	UpdatedAt time.Time
}

func (o *Order) Add() error {
	return GetDB().Create(o).Error
}
