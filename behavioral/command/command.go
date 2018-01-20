package command

import (
	"fmt"
)

type Order interface {
	execute()
}

type Stock struct {
	name string
	count int
}
func (s *Stock) Buy()  {
	fmt.Println("Bought new item.")
	s.count += 1
}
func (s *Stock) Sell()  {
	fmt.Println("Sold one item.")
	s.count -= 1
}

type BuyStock struct {
	stock Stock
}
func (b *BuyStock) execute()  {
	b.stock.Buy()
}

type SellStock struct {
	stock Stock
}
func (b *SellStock) execute()  {
	b.stock.Sell()
}

type Broker struct {
	orders []Order
}
func (b *Broker) TakeOrder(o Order)  {
	b.orders = append(b.orders, o)
}
func (b *Broker) PlaceOrders()  {
	for _,v := range b.orders {
		v.execute()
	}

	b.orders = make([]Order, 0)
}

func Demo() {
	stock := Stock{name:"phone", count:2}
	buyStck := &BuyStock{stock}
	sellStck := &SellStock{stock}
	broker := &Broker{}
	broker.orders = make([]Order, 0)

	broker.TakeOrder(buyStck)
	broker.TakeOrder(sellStck)

	broker.PlaceOrders()
	
}