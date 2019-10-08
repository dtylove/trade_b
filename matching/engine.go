package matching

import (
	"github.com/shopspring/decimal"
)

const asks = "asks"
const bids = "bids"

type Engine struct {
	MarketId         int
	OrderBookManager BookManager
	//Traded           chan Offer
	Canceled chan int
	//Options          Options
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

func (engine *Engine) AskOrderBook() OrderBook {
	return engine.OrderBookManager.AskOrderBook
}

func (engine *Engine) BidOrderBook() OrderBook {
	return engine.OrderBookManager.BidOrderBook
}

func (engine *Engine) Submit(order *Order) {
	bidsBook, askBook := engine.OrderBookManager.GetBooks()
	if order.OrderType == asks {
		engine.match(order, bidsBook)
		if order.IsRemain {
			askBook.Add(order)
		}
	} else if order.OrderType == bids{
		engine.match(order, askBook)
		if order.IsRemain {
			bidsBook.Add(order)
		}
	}
}

func (engine *Engine) Cancel(order *Order) {
	book, _ := engine.OrderBookManager.GetBooks()
	book.Remove(order)
}

func (engine *Engine) match(order *Order, orderBook OrderBook) {
	iter := orderBook.Book.Iterator()

	for iter.Next() && order.IsRemain {
		tmpPrice := iter.Key()
		curPrice, _ := decimal.NewFromString(tmpPrice.(string))
		if order.IsCrossed(curPrice) {
			pContainer := iter.Value().(PriceContainer)
			orderSlice := pContainer.Orders
			for _, o := range orderSlice {
				order.TradeWith(o)
				if !order.IsRemain {
					break
				}
				if !o.IsRemain {
					// TODO 广播?
					orderBook.Remove(o)
				}
			}
		}
	}
}
