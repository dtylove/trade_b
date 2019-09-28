package matching

import "github.com/shopspring/decimal"

// 相同价格的挂单放到数组内 尽量减轻树频繁增删改查的压力
type PriceContainer struct {
	Price  decimal.Decimal
	Orders []Order
}

func InitContainer(price decimal.Decimal) (pl PriceContainer) {
	pl.Price = price
	pl.Orders = []Order{}
	return
}

func (pl *PriceContainer) Top() Order {
	return pl.Orders[0]
}

func (pl *PriceContainer) IsEmpty() (result bool) {
	if len(pl.Orders) == 0 {
		result = true
	}
	return
}

func (pl *PriceContainer) Add(order Order) {
	pl.Orders = append(pl.Orders, order)
	return
}

func (pl *PriceContainer) Remove(order Order) {
	var orders []Order
	for index, o := range pl.Orders {
		if o.Id == order.Id {
			pl.Orders = append(orders[index:], orders[:index +1]...)
			break
		}
	}
	return
}

func (pl *PriceContainer) Find(id int64) (order Order) {
	for _, o := range pl.Orders {
		if o.Id == id {
			order = o
			break
		}
	}
	return
}
