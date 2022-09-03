package hw04lrucache

type List interface {
	Len() int
	Front() *ListItem
	Back() *ListItem
	PushFront(v interface{}) *ListItem
	PushBack(v interface{}) *ListItem
	Remove(i *ListItem)
	MoveToFront(i *ListItem)
}

type ListItem struct {
	Value interface{}
	Next  *ListItem
	Prev  *ListItem
}

type list struct {
	Head   *ListItem
	Tail   *ListItem
	length int
}

func (l *list) Len() int {
	return l.length
}

func (l *list) Front() *ListItem {
	return l.Head
}

func (l *list) Back() *ListItem {
	return l.Tail
}

func (l *list) PushFront(v interface{}) *ListItem {
	node := &ListItem{Value: v}
	if l.Head == nil {
		l.Head = node
		l.Tail = node
	} else {
		node.Next = l.Head
		l.Head.Prev = node
		l.Head = node
	}
	l.length++
	return node
}

func (l *list) PushBack(v interface{}) *ListItem {
	node := &ListItem{Value: v}
	if l.Tail == nil {
		l.Head = node
		l.Tail = node
	} else {
		l.Tail.Next = node
		node.Prev = l.Tail
		l.Tail = node
	}
	l.length++
	return node
}

func (l *list) Remove(i *ListItem) {
	switch i {
	case l.Head:
		i.Next.Prev = nil
		l.Head = i.Next
	case l.Tail:
		l.Tail = i.Prev
		l.Tail.Next = nil
		i.Prev = nil
	default:
		i.Next.Prev = i.Prev
		i.Prev.Next = i.Next

		i.Next = nil
		i.Prev = nil
	}
	l.length--
}

func (l *list) MoveToFront(i *ListItem) {
	l.Remove(i)
	l.Head.Prev = i
	i.Next = l.Head
	l.Head = i
}

func NewList() List {
	return &list{
		Head:   nil,
		Tail:   nil,
		length: 0,
	}
}
