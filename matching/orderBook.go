package matching

import (
	"encoding/json"
	"fmt"
	rbt "github.com/emirpasic/gods/trees/redblacktree"
	"github.com/pkg/errors"
)

type OrderBook struct {
	MarketId  int       // 市场id
	Side      string    // asks 或 bids
	Book      *rbt.Tree // 市场中的挂单
	Broadcast bool      // 广播
}

func InitOrderBook(marketId int, side string) (ob OrderBook) {
	ob.MarketId = marketId
	ob.Side = side
	ob.Book = rbt.NewWithStringComparator()
	return
}

func (ob *OrderBook) Find(order Order) (o Order, err error) {
	if ob.Book.Size() == 0 {
		err = errors.New("book中没有order")
		return
	}

	values, _ := ob.Book.Get(order.Price.String())
	if values == nil {
		err = errors.New("找不到该订单")
		return
	}
	priceContainer := values.(PriceContainer)
	o = priceContainer.Find(order.Id)

	return
}

func (ob *OrderBook) Add(order Order) (err error) {
	if ob.Side != order.OrderType {
		return errors.New("book类型 与 订单类型不匹配 book.Side: " +
			ob.Side + " orderType: " + order.OrderType)
	}

	if ob.Side == Asks && order.IsBuy == true {
		return errors.New("book类型为 Asks order.IsBuy 必须为 false ")
	}

	if ob.Side == Bids && order.IsBuy == false {
		return errors.New("book类型为 Bids order.IsBuy 必须为 true ")
	}

	container := InitContainer(order.Price)
	values, found := ob.Book.Get(order.Price.String())
	if found {
		container = values.(PriceContainer)
	}
	container.Add(order)
	ob.Book.Put(order.Price.String(), container)

	return
}

// 删除必须使用 价格 和 order id
func (ob *OrderBook) Remove(order Order) (o Order) {
	if ob.Book.Size() == 0 {
		return order
	}
	values, found := ob.Book.Get(order.Price.String())
	if !found {
		return
	}

	priceContainer := values.(PriceContainer)
	o = priceContainer.Find(order.Id)
	if o.Id == 0 {
		return
	}

	priceContainer.Remove(order.Id)
	if !priceContainer.IsEmpty() {
		ob.Book.Remove(order.Price.String())
	}

	return
}

func (ob *OrderBook) Size() int {
	return ob.Book.Size()
}

func (ob *OrderBook) Top() (order *Order) {
	var container PriceContainer
	if ob.Side == "asks" {
		container = ob.Book.Left().Value.(PriceContainer)
	} else {
		container = ob.Book.Right().Value.(PriceContainer)
	}

	return container.Top()
}

func (ob *OrderBook) Print(max int) {
	//fmt.Println(ob.Book)
	iter := ob.Book.Iterator()
	count := 0
	for ; iter.Next() && count < max; {
		fmt.Println("TreeKey(price): ", iter.Key())

		container := iter.Value().(PriceContainer)
		fmt.Println("ContainerPrice: ", container.Price)
		//fmt.Println("OrdersContainer:  ", container.Orders)
		orderSilce := container.Orders
		for _, v := range orderSilce {
			orderJson, _ := json.Marshal(v)
			fmt.Println("Order:  ", string(orderJson))
		}
		fmt.Print("\n")
		count++
	}

}
