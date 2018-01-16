package memento

import (
	"errors"
)

type State struct {
	Description string
}

type memento struct {
	state State
}

type originator struct {
	state State
}
func (o *originator) StoreStateToMemento() memento {
	return memento{state: o.state}
}
func (o *originator) StoreStateFromMemento(m memento) {
	o.state = m.state
}

type careTaker struct {
	mementos []memento
}
func (c *careTaker) Push(m memento)  {
	c.mementos = append(c.mementos, m)

}
func (c *careTaker) Get(i int) (memento, error)  {
	var m memento
	if len(c.mementos) < i || i< 0 {
		return m, errors.New("Index is out of range")
	}
	
	return c.mementos[i], nil
}