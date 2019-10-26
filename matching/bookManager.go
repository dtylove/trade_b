package matching

type BookManager struct {
	MarketId     uint
	AsksOrderBook OrderBook
	BidsOrderBook OrderBook
}
/**
 * param  marketId 市场id 不同货币对 对应不同的id
 */
func InitBookManager(marketId uint) (orderBookManager BookManager) {
	orderBookManager.MarketId = marketId
	orderBookManager.AsksOrderBook = InitOrderBook(marketId, Asks)
	orderBookManager.BidsOrderBook = InitOrderBook(marketId, Bids)
	return
}

func (obm *BookManager) GetBooks() (OrderBook, OrderBook) {

	return obm.AsksOrderBook, obm.BidsOrderBook

}
