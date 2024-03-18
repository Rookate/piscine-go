package piscine

func ListSize(l *List) int {
	wagon := l.Head
	count := 0

	for wagon != nil {
		wagon = wagon.Next
		count++
	}
	return count
}
