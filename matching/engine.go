package matching

type Engine struct {
	MarketId    int
	BookManager BookManager

}

func InitEngine(marketId int) Engine {
	return Engine{
		MarketId:    marketId,
		BookManager: InitBookManager(marketId),
	}
}

func (engine *Engine) AskOrderBook() OrderBook {
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
