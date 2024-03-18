package piscine

type NodeL struct {
	Data interface{}
	Next *NodeL
}

type List struct {
	Head *NodeL
	Tail *NodeL
}

func ListSize(l *List) int {
	wagon := l.Head
	count := 0 

	for _, wagon != nil {
		wagon = wagon.Next
		count ++
	}
	return count
}
