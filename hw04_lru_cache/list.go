package hw04lrucache

import (
	"fmt"
)

type List interface {
	Len() int
	Front() *ListItem
	Back() *ListItem
	PushFront(v interface{}) *ListItem
	PushBack(v interface{}) *ListItem
	Remove(i *ListItem)
	MoveToFront(i *ListItem)
	printList()
}

type ListItem struct {
	Value interface{}
	Next  *ListItem
	Prev  *ListItem
}

func initList(l *list, value interface{}) {
	item := ListItem{Next: nil, Prev: nil, Value: value}
	l.head, l.tail = &item, &item
	l.elem_count++
}

func (l *list) MoveToFront(i *ListItem) {
	if i == l.head {
		return
	}
	if i.Prev != nil {
		i.Prev.Next = i.Next
	}
	if i.Next != nil {
		i.Next.Prev = i.Prev
	}
	i.Prev = nil
	i.Next = l.head
	l.head.Prev = i
	l.head = i
	//TODO
}

func (l *list) Remove(i *ListItem) {
	i.Prev.Next = i.Next
	i.Next.Prev = i.Prev
	l.elem_count--
}

func (l *list) PushFront(v interface{}) *ListItem {
	if l.head == nil {
		initList(l, v)
		return l.head
	}
	var lItem = ListItem{Value: v, Prev: nil, Next: l.head}
	l.head.Prev = &lItem
	l.head = &lItem
	l.elem_count++
	return &lItem
}

func (l *list) PushBack(v interface{}) *ListItem {
	if l.tail == nil {
		initList(l, v)
		return l.tail
	}
	var lItem = ListItem{Value: v, Prev: l.tail, Next: nil}
	l.tail.Next = &lItem
	l.tail = &lItem
	l.elem_count++
	return &lItem
}

func (l *list) Front() *ListItem {
	if l.head == nil {
		return nil
	}
	return l.head
}

func (l *list) Back() *ListItem {
	if l.tail == nil {
		return nil
	}
	return l.tail
}

func (l *list) Len() int {
	return l.elem_count
}

type list struct {
	//List       // Remove me after realization.
	elem_count int
	head       *ListItem
	tail       *ListItem
	// Place your code here.
}

func (l *list) printList() {
	currentElem := l.head
	for currentElem != nil {
		fmt.Println(currentElem.Value)
		currentElem = currentElem.Next
	}
}

// removeByIndex удаляет элемент по индексу
func removeByIndex(l *list, index int) {
	currentElem := l.head
	for i := 0; currentElem != nil; i++ {
		if i == index {
			l.MoveToFront(currentElem)
		}
		currentElem = currentElem.Next
	}
}

func NewList() List {
	return new(list)
}
