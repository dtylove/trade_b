package main

import (
	"github.com/shopspring/decimal"
	"mathcing/matching"
)

func main() {

	var order matching.Order
	order.Id = 1                                   // 订单id
	order.IsBuy = false                            // 买true bids 卖false asks
	order.BookId = 1                               // order book id (针对不同市场 货币对)
	order.OrderType = "asks"                       // asks bids
	order.Price, _ = decimal.NewFromString("2.2")    // 价格
	order.Quantity, _ = decimal.NewFromString("5") // 总个数
	order.Remained = order.Quantity
	order.UserId = 1
	order.Id = 1
	order.IsRemain = true

	var order2 matching.Order
	order2.Id = 2                                   // 订单id
	order2.IsBuy = true                             // 买true bids 卖false asks
	order2.BookId = 1                               // order book id (针对不同市场 货币对)
	order2.OrderType = "bids"                       // asks bids
	order2.Price, _ = decimal.NewFromString("2.1")    // 价格
	order2.Quantity, _ = decimal.NewFromString("2") // 总个数
	order2.Remained = order2.Quantity
	order2.UserId = 1
	order2.IsRemain = true

	var order3 matching.Order
	order3.Id = 3                                   // 订单id
	order3.IsBuy = true                             // 买true bids 卖false asks
	order3.BookId = 1                               // order book id (针对不同市场 货币对)
	order3.OrderType = "bids"                       // asks bids
	order3.Price, _ = decimal.NewFromString("1")    // 价格
	order3.Quantity, _ = decimal.NewFromString("2.9") // 总个数
	order3.Remained = order3.Quantity
	order3.UserId = 1
	order3.IsRemain = true

	engine := matching.InitEngine(1)
	engine.Submit(order)
	engine.Submit(order2)
	//engine.Submit(order3)
	abook, bbook := engine.OrderBookManager.GetBooks()
	abook.Print(10)
	bbook.Print(10)

}
