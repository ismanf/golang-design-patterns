package main

import (
	"fmt"
)

//Iterator object should implement this interface
type Iterator interface {
	First()
	Next() bool
	Done() bool
	CurrentItem() interface{}
}

//A list object should implement Iterable interface in order to be a iterable list
type Iterable interface {
	CreateIterator() Iterator
}

//Iterator object
type ListIterator struct {
	data  []string
	index int
}

func (i *ListIterator) Done() bool {
	return i.index >= len(i.data)
}

func (i *ListIterator) CurrentItem() interface{} {
	return i.data[i.index]
}

func (i *ListIterator) First() {
	i.index = 0
}

func (i *ListIterator) Next() bool {
	i.index++
	if i.Done() {
		return false
	}
	return true
}

//Actual a basic list object.I don't use abstraction here.
type List struct {
	data []string
}

func (l *List) Append(str string) {
	l.data = append(l.data, str)
}

func (l *List) Count() int {
	return len(l.data)
}

func (l *List) CreateIterator() Iterator {
	listIterator := &ListIterator{data: l.data, index: -1}
	return listIterator
}

//Basic testing....

func main() {
	lst := &List{data: make([]string, 0)}
	lst.Append("A")
	lst.Append("B")
	lst.Append("C")
	lst.Append("D")
	lst.Append("E")

	//time to iterate...
	lstIterator := lst.CreateIterator()
	for lstIterator.Next() {
		fmt.Println("-", lstIterator.CurrentItem())
	}

	lstIterator.First()
	fmt.Println("First:", lstIterator.CurrentItem())
}
