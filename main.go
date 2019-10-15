package main

import (
	"github.com/shopspring/decimal"
	"mathcing/matching"
)

func main() {
	bm := matching.InitBookManager(1)

	var order matching.Order
	order.Id = 1                                   // 订单id
	order.IsBuy = false                            // 买true bids 卖false asks
	order.BookId = 1                               // order book id (针对不同市场 货币对)
	order.OrderType = "asks"                       // asks bids
	order.Price, _ = decimal.NewFromString("1")    // 价格
	order.Volume, _ = decimal.NewFromString("1")   // 总价值 总个数*价格
	order.Quantity, _ = decimal.NewFromString("1") // 总个数
	order.UserId = 1
	order.Id = 1

	abook, _ := bm.GetBooks()

	err := abook.Add(order)
	if err != nil {
		panic(err)
	}
	//bbook.Add(order)
	//abook.Print(10)
	//bbook.Print(10)
	o, err := abook.Top()
	if err != nil {
		panic(err)
	}
	o.Print()
}
