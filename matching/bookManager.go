package matching

type BookManager struct {
	MarketId     int
	AskOrderBook OrderBook
	BidOrderBook OrderBook
}
/**
 * param  marketId 市场id 不同货币对 对应不同的id
 */
func InitBookManager(marketId int) (orderBookManager BookManager) {
	orderBookManager.MarketId = marketId
	orderBookManager.AskOrderBook = InitOrderBook(marketId, Asks)
	orderBookManager.BidOrderBook = InitOrderBook(marketId, Bids)
	return
}

func (obm *BookManager) GetBooks() (OrderBook, OrderBook) {

	return obm.AskOrderBook, obm.BidOrderBook

}
