package matching

import (
	"encoding/json"
	"fmt"
	"github.com/shopspring/decimal"
)

type Engine struct {
	TxId             uint       // 交易id
	MarketId         uint         // 市场id
	OrderBookManager BookManager // book handler

	//Options          Options
	Canceled  chan int
	orderPipe chan Order // 负责从各种消息中间件中获取order
}

type CounterParty struct {
	OrderId uint           // 订单号
	Pricing decimal.Decimal // 定价
	MatchQ  decimal.Decimal // 成交量
	Funds   decimal.Decimal // 成交价钱
}

// taker 购买者 maker 出票人 卖出人
type Trade struct {
	TxId     uint           // 交易id
	MarketId uint           // 市场id
	Price    decimal.Decimal // 成交价
	MatchQ   decimal.Decimal // 成交量
	Funds    decimal.Decimal // 成交总额
	Taker    CounterParty    // 购买者
	Maker    CounterParty    // 出货者
}

func InitEngine(marketId uint) *Engine {
	return &Engine{
		MarketId:         marketId,
		OrderBookManager: InitBookManager(marketId),
		//Options:          options,
		//Traded:           make(chan Offer, 5),
		Canceled:  make(chan int, 3),
		orderPipe: make(chan Order, 1024),
	}
}

func (engine *Engine) AsksOrderBook() OrderBook {
	return engine.OrderBookManager.AsksOrderBook
}

func (engine *Engine) BidsOrderBook() OrderBook {
	return engine.OrderBookManager.BidsOrderBook
}

// order的入口
func (engine *Engine) Submit(order Order) {
	engine.match(&order)

	if order.IsRemain {
		engine.Add(order)
	}
}

// 取消用
func (engine *Engine) Cancel(order Order) {
	var orderBook OrderBook
	if order.OrderType == Asks {
		orderBook = engine.AsksOrderBook()
	} else {
		orderBook = engine.BidsOrderBook()
	}

	orderBook.Remove(order)
}

// 完全成交用
func (engine *Engine) Remove(order Order) {
	var orderBook OrderBook
	if order.OrderType == Asks {
		orderBook = engine.AsksOrderBook()
	} else {
		orderBook = engine.BidsOrderBook()
	}

	orderBook.Remove(order)
}

func (engine *Engine) match(order *Order) {
	var orderBook OrderBook
	if order.OrderType == Asks {
		orderBook = engine.BidsOrderBook()
	} else {
		orderBook = engine.AsksOrderBook()
	}

	if orderBook.Size() == 0 {
		return
	}

	marketOrder := orderBook.Top()

	if order.IsCrossed(marketOrder.Price) {
		tradeInfo := order.TradeWith(marketOrder)
		data, _ := json.Marshal(tradeInfo)
		fmt.Println(string(data))
		// TODO 通知 marketOrder

		if marketOrder.IsRemain == false {
			engine.Remove(*marketOrder)
		}
	}
}

func (engine *Engine) Add(order Order) {
	var orderBook OrderBook
	if order.OrderType == Asks {
		orderBook = engine.AsksOrderBook()
	} else {
		orderBook = engine.BidsOrderBook()
	}

	err := orderBook.Add(order)
	if err != nil {
		fmt.Println(err)
	}
}

// 不断接收order
func (engine *Engine) RunOrderApplier() {
	for {
		select {
		case order := <-engine.orderPipe:
			engine.Submit(order)
		}
	}
}

func (engine *Engine) RunOrderFetcher() {
	//var order Order
	//order.Id = 1                                   // 订单id
	//order.IsBuy = false                            // 买true bids 卖false asks
	//order.BookId = 1                               // order book id (针对不同市场 货币对)
	//order.OrderType = "asks"                       // asks bids
	//order.Price, _ = decimal.NewFromString("2.2")  // 价格
	//order.Quantity, _ = decimal.NewFromString("5") // 总个数
	//order.Remained = order.Quantity
	//order.UserId = 1
	//order.Id = 1
	//order.IsRemain = true
	//
	//var order2 Order
	//order2.Id = 2                                   // 订单id
	//order2.IsBuy = true                             // 买true bids 卖false asks
	//order2.BookId = 1                               // order book id (针对不同市场 货币对)
	//order2.OrderType = "bids"                       // asks bids
	//order2.Price, _ = decimal.NewFromString("2.1")  // 价格
	//order2.Quantity, _ = decimal.NewFromString("2") // 总个数
	//order2.Remained = order2.Quantity
	//order2.UserId = 1
	//order2.IsRemain = true
	//
	//var order3 Order
	//order3.Id = 3                                     // 订单id
	//order3.IsBuy = true                               // 买true bids 卖false asks
	//order3.BookId = 1                                 // order book id (针对不同市场 货币对)
	//order3.OrderType = "bids"                         // asks bids
	//order3.Price, _ = decimal.NewFromString("1")      // 价格
	//order3.Quantity, _ = decimal.NewFromString("2.9") // 总个数
	//order3.Remained = order3.Quantity
	//order3.UserId = 1
	//order3.IsRemain = true
	////
	//engine.orderPipe <- order
	//engine.orderPipe <- order2
	//engine.orderPipe <- order3

}
