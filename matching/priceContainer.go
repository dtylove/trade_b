package matching

import (
	"github.com/shopspring/decimal"
)

// 相同价格的挂单放到数组内
type PriceContainer struct {
	Price  decimal.Decimal
	Orders []Order
	TotalQ decimal.Decimal
}

func InitContainer(price decimal.Decimal) *PriceContainer {
	return &PriceContainer{
		Price:  price,
		Orders: []Order{},
	}
}

func (pl *PriceContainer) Top() *Order {
	return &pl.Orders[0]
}

func (pl *PriceContainer) IsEmpty() (result bool) {
	if len(pl.Orders) == 0 {
		result = true
	}
	return
}

// isIncr true 增长 false 减少
func (pl *PriceContainer) ChangeTotalQ(quantity decimal.Decimal, isIncr bool) {
	if isIncr {
		pl.TotalQ = pl.TotalQ.Add(quantity)
	} else {
		pl.TotalQ = pl.TotalQ.Sub(quantity)
	}
}

func (pl *PriceContainer) Add(order Order) {
	if len(pl.Orders) == 0 {
		pl.TotalQ = order.Remained
	} else {
		pl.ChangeTotalQ(order.Remained, true)
	}
	pl.Orders = append(pl.Orders, order)
	return
}

func (pl *PriceContainer) Remove(orderId uint) {
	for index, o := range pl.Orders {
		if o.Id == orderId {
			pl.Orders = append(pl.Orders[index:], pl.Orders[:index+1]...)
			break
		}
	}
	return
}

func (pl *PriceContainer) Find(id uint) (order Order) {
	for _, o := range pl.Orders {
		if o.Id == id {
			order = o
			break
		}
	}
	return
}
