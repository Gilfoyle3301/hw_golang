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

func NewList() List {
	return new(list)
}

type list struct {
	length       int
	beginElement *ListItem
	endElement   *ListItem
}

func (l *list) Len() int {
	return l.length
}

func (l *list) Front() *ListItem {
	return l.beginElement
}

func (l *list) Back() *ListItem {
	return l.endElement
}

func (l *list) PushFront(v interface{}) *ListItem {
	mod := &ListItem{Value: v, Next: l.beginElement, Prev: nil}

	if l.Len() == 0 {
		l.beginElement = mod
		l.endElement = mod
		l.length++
		return l.beginElement
	}

	if l.Len() != 0 {
		l.beginElement.Prev = mod
		l.beginElement = mod
		l.length++
	}
	return l.beginElement
}

func (l *list) PushBack(v interface{}) *ListItem {
	mod := &ListItem{Value: v, Next: nil, Prev: l.endElement}

	if l.Len() == 0 {
		l.beginElement = mod
		l.endElement = mod
		l.length++
		return l.endElement
	}

	if l.Len() != 0 {
		l.endElement.Next = mod
		l.endElement = mod
		l.length++
	}
	return l.endElement
}

func (l *list) Remove(i *ListItem) {
	switch {
	case i.Prev == nil:
		l.beginElement = l.beginElement.Next
		l.beginElement.Prev = nil
		l.length--
	case i.Next == nil:
		l.endElement = l.endElement.Prev
		l.endElement.Next = nil
		l.length--
	default:
		i.Prev.Next = i.Next
		i.Next.Prev = i.Prev
		l.length--
	}
}

func (l *list) MoveToFront(i *ListItem) {
	switch {
	case i.Prev == nil:
		return
	case i.Next == nil:
		l.endElement = l.endElement.Prev
		l.endElement.Next = nil
	default:
		i.Next.Prev = i.Prev
		i.Prev.Next = i.Next
	}
	i.Next = l.beginElement
	i.Prev = nil
	l.beginElement = i
}
