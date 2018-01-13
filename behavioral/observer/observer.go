package observer

import (
	"fmt"
)

type Observer interface {
	Notify(string)
}

type Publisher struct {
	ObserverList []Observer
}

func (p *Publisher) Subscribe(o Observer) {
	p.ObserverList = append(p.ObserverList, o)
}

func (p *Publisher) Unsubscribe(o Observer) {
	var index int
	for i, v := range p.ObserverList {
		if v == o {
			index = i
			break
		}
	}

	p.ObserverList = append(p.ObserverList[:index], p.ObserverList[index + 1:]...)
}

func (p *Publisher) Publish(m string) {
	fmt.Println("Got new message! Publishing...")
	for _, o := range p.ObserverList {
		o.Notify(m)
	}
	fmt.Println("Puplished!")
}

