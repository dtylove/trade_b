package matching

import (
	rbt "github.com/emirpasic/gods/trees/redblacktree"
)

type OrderBook struct {
	MarketId  int       // 市场id
	Side      string    // asks 或 bids
	Book      *rbt.Tree // 市场中的挂单
	Broadcast bool      // 广播
}

func InitOrderBook(marketId int, side string) (ob OrderBook) {
	ob.MarketId = marketId
	ob.Side = side
	ob.Book = rbt.NewWithStringComparator()
	return
}

func (ob *OrderBook) Find(order Order) (o Order) {
	values, _ := ob.Book.Get(order.Price.String())
	priceLevel := values.(PriceContainer)
	o = priceLevel.Find(order.Id)
	return
}

func (ob *OrderBook) Add(order Order) error {
	container := InitContainer(order.Price)
	values, found := ob.Book.Get(order.Price.String())
	if found {
		container = values.(PriceContainer)
	}
	container.Add(order)
	ob.Book.Put(order.Price.String(), container)

	return nil
}

func (ob *OrderBook) Remove(order Order) (o Order) {
	values, found := ob.Book.Get(order.Price.String())
	if !found {
		return
	}
	priceLevel := values.(PriceContainer)
	o = priceLevel.Find(order.Id)
	if o.Id == 0 {
		return
	}
	priceLevel.Remove(order)
	if priceLevel.IsEmpty() {
		ob.Book.Remove(order.Price.String())
	}

	return
}
