package models

import (
	"dtyTrade/matching"
	"time"
)

type Order struct {
	matching.Order // 提交到市场的order
	CreatedAt time.Time
	UpdatedAt time.Time
}

func (o *Order) Add() error {
	return GetDB().Create(o).Error
}
