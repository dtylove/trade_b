package matching

<<<<<<< HEAD
type Engine struct {
	MarketId    int
	BookManager BookManager

=======
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
>>>>>>> c4c5ffd9075d29e35c01729bfe027fd2ab01b1ee
}

func InitEngine(marketId int) Engine {
	return Engine{
<<<<<<< HEAD
		MarketId:    marketId,
		BookManager: InitBookManager(marketId),
=======
		MarketId:         marketId,
		OrderBookManager: InitBookManager(marketId),
		//Options:          options,
		//Traded:           make(chan Offer, 5),
		Canceled: make(chan int, 3),
>>>>>>> c4c5ffd9075d29e35c01729bfe027fd2ab01b1ee
	}
}

func (engine *Engine) AskOrderBook() OrderBook {
<<<<<<< HEAD
	return engine.BookManager.AskOrderBook
}

func (engine *Engine) BidOrderBook() OrderBook {
	return engine.BookManager.BidOrderBook
}

//func (engine *Engine) Submit(order Order) {
//	book, counterBook := engine.BookManager.GetBooks(order.OrderType)
//	engine.match(order, counterBook)
//	engine.addOrCancel(order, book)
//}
//
//func (engine *Engine) Cancel(order Order) {
//	book, _ := engine.BookManager.GetBooks(order.OrderType)
//	rOrder := book.Remove(order)
//	if rOrder.Id != 0 {
//		engine.publishCancel(rOrder, "cancelled by user")
//	} else {
//		// Matching.logger.warn "Cannot find order##{order.id} to cancel, skip."
//	}
//
//}

//func (engine *Engine) match(order Order, counterBook OrderBook) {
//	if order.IsFilled() || engine.isTiny(order) {
//		return
//	}
//	counterOrder := counterBook.Top()
//	if counterOrder.Id == 0 {
//		return
//	}
//	trade := order.TradeWith(counterOrder, counterBook)
//	if trade.isNotValidated() {
//		return
//	}
//	counterBook.FillTop(trade)
//	order.Fill(trade)
//	engine.publish(order, counterOrder, trade)
//	engine.match(order, counterBook)
//}
//
//func (engine *Engine) addOrCancel(order Order, book OrderBook) {
//	if order.IsFilled() {
//		return
//	}
//	if order.OrderType == "LimitOrder" {
//		book.Add(order)
//	} else if order.OrderType == "MarketOrder" {
//		engine.publishCancel(order, "fill or kill market order")
//	}
//	return
//}
//
//func (engine *Engine) publish(order, counterOrder Order, trade Trade) {
//	var ask, bid Order
//	if order.Type == "ask" {
//		ask = order
//		bid = counterOrder
//	} else {
//		ask = counterOrder
//		bid = order
//	}
//	offer := Offer{
//		MarketId:    order.MarketId,
//		AskId:       ask.Id,
//		BidId:       bid.Id,
//		StrikePrice: trade.Price,
//		Volume:      trade.Volume,
//		Funds:       trade.Funds,
//	}
//	engine.Traded <- offer
//	// logger 记录订单成交
//	return
//}
//
//func (engine *Engine) publishCancel(order Order, reson string) {
//	engine.Canceled <- order.Id
//	// logger 记录订单取消
//	return
//}
//
//func (engine *Engine) isTiny(order Order) (result bool) {
//	var fixed = DEFAULT_PRECISION
//	if engine.Options.Ask.Fixed != 0 {
//		fixed = engine.Options.Ask.Fixed
//	}
//	cas := decimal.NewFromFloat(1)
//	for fixed > 0 {
//		cas = cas.Mul(decimal.NewFromFloat(10))
//		fixed--
//	}
//	minVolume := decimal.NewFromFloat(1.0).Div(cas)
//	return order.Volume.LessThan(minVolume)
//}
//
//func (engine *Engine) LimitOrders() (result map[string]map[string][]Order) {
//	askOrderBook := engine.AskOrderBook()
//	bidOrderBook := engine.BidOrderBook()
//	result["ask"] = askOrderBook.LimitOrdersMap()
//	result["bid"] = bidOrderBook.LimitOrdersMap()
//	return
//}
//
//func (engine *Engine) MarketOrders() (result map[string][]Order) {
//	askOrderBook := engine.AskOrderBook()
//	bidOrderBook := engine.BidOrderBook()
//	result["ask"] = askOrderBook.MarketOrdersMap()
//	result["bid"] = bidOrderBook.MarketOrdersMap()
//	// result["ask"] = engine.AskOrderBook.MarketOrders()
//	// result["bid"] = engine.BidOrderBook.MarketOrders()
//	return
//}
=======
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
>>>>>>> c4c5ffd9075d29e35c01729bfe027fd2ab01b1ee
