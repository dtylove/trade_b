package matching

import (
	"encoding/json"
	"fmt"
	"github.com/shopspring/decimal"
)

type Engine struct {
	TxId             int64       // 交易id
	MarketId         int         // 市场id
	OrderBookManager BookManager // book handler
	//Traded           chan Offer
	Canceled chan int
	//Options          Options
}

type CounterParty struct {
	OrderId int64           // 订单号
	Pricing decimal.Decimal // 定价
	MatchQ  decimal.Decimal // 成交量
	Funds   decimal.Decimal // 成交价钱
}

// taker 购买者 maker 出票人 卖出人
type Trade struct {
	TxId     int64           // 交易id
	MarketId int64           // 市场id
	Price    decimal.Decimal // 成交价
	MatchQ   decimal.Decimal // 成交量
	Funds    decimal.Decimal // 成交总额
	Taker    CounterParty    // 购买者
	Maker    CounterParty    // 出货者
}

func InitEngine(marketId int) Engine {
	return Engine{
		MarketId:         marketId,
		OrderBookManager: InitBookManager(marketId),
		//Options:          options,
		//Traded:           make(chan Offer, 5),
		Canceled: make(chan int, 3),
	}
}

func (engine *Engine) AsksOrderBook() OrderBook {
	return engine.OrderBookManager.AsksOrderBook
}

func (engine *Engine) BidsOrderBook() OrderBook {
	return engine.OrderBookManager.BidsOrderBook
}

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
