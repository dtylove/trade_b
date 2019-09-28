package order

import rbt "github.com/emirpasic/gods/trees/redblacktree"

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
