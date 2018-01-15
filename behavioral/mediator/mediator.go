package mediator

import (
	"fmt"
	"errors"
)

// Mediator
type TradeService struct {
	Orders []Order
}

func (t *TradeService) GetOrders(orderType string) ([]Order, error) {
	orders := make([]Order, 0)

	if orderType != "buy" || orderType != "sell" {
		return orders, errors.New("Invalid order type provided.")
	}

	for _, v := range t.Orders {
		if v.OrderType == orderType {
			orders = append(orders, v)
		}
	}

	return orders, nil
}

func (t *TradeService) PushOrder(o Order)  {
	// We dont need too much logic to show the purpose of the pattern
	t.Orders = append(t.Orders, o)
}

type Order struct {
	OrderType string
	Price float32
}

type Trader struct {
	Service *TradeService
	OrderToPlace Order
	InterestedOrders []Order
}
func (t *Trader) PlaceOrder()  {
	t.Service.PushOrder(t.OrderToPlace)
}
func (t *Trader) GetOrders(orderType string)  {
	t.InterestedOrders, _ = t.Service.GetOrders(orderType)
}

func Demo ()  {
	tradeService := &TradeService{}
	tradeService.Orders = make([]Order, 0)

	traderOne := &Trader{ Service: tradeService}
	traderOne.OrderToPlace = Order{ "buy", 1.2 }
	traderOne.PlaceOrder()

	traderTwo := &Trader{ Service: tradeService}
	traderTwo.OrderToPlace = Order{ "sell", 1.5 }
	traderOne.PlaceOrder()

	traderOne.GetOrders("sell")
	fmt.Println(traderOne.InterestedOrders)

	traderTwo.GetOrders("buy")
	fmt.Println(traderTwo.InterestedOrders)
}
