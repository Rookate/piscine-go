package piscine

func ListPushFront(l *List, data interface{}) {
	wagon := &NodeL{Data: data}
	if l.Head == nil {
		l.Head = wagon
		return
	}
	wagon.Next = l.Head
	l.Head = wagon
}
