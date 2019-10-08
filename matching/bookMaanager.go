package matching

type BookManager struct {
	MarketId     int
	AskOrderBook OrderBook
	BidOrderBook OrderBook
}

func InitBookManager(marketId int) (orderBookManager BookManager) {
	orderBookManager.MarketId = marketId
	orderBookManager.AskOrderBook = InitOrderBook(marketId, "ask")
	orderBookManager.BidOrderBook = InitOrderBook(marketId, "bid")
	return
}

func (obm *BookManager) GetBooks() (OrderBook, OrderBook) {

	return obm.AskOrderBook, obm.BidOrderBook

}
