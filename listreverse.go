package piscine

func ListReverse(l *List) {
	var next, prev *NodeL
	current := l.Head
	l.Tail = l.Head

	for current != nil {
		next = current.Next
		current.Next, prev, current = prev, current, next
	}
	l.Head, l.Tail = prev, l.Head
}
