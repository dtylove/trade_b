package models

import (
	"fmt"
	"github.com/gin-gonic/gin/json"
	"time"
)

type Product struct {
	Id        uint `gorm:"column:id;primary_key"`
	CreatedAt time.Time
	UpdatedAt time.Time

	Base      string                              // 买方货币 BTC  ETH等
	BaseMinQ  string `sql:"type:decimal(32,16);"` // base 最小挂单量
	BaseMaxQ  string `sql:"type:decimal(32,16);"` // 最大
	BaseScale int32                               // 最大小数位

	Counter      string                              // 卖方货币 BTC 等
	CounterMinQ  string `sql:"type:decimal(32,16);"` // counter最小挂单量
	CounterMaxQ  string `sql:"type:decimal(32,16);"` // 最大
	CounterScale int32                               // 最大小数位
}

func (p *Product) Add() error {
	data, _ := json.Marshal(p)
	fmt.Println(string(data))
	return GetDB().Create(p).Error
}

func (p *Product) FindById() error {
	return GetDB().Find(p, p.Id).Error
}
