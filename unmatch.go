package piscine

func Unmatch(a []int) int {
	numCount := make(map[int]int)

	for _, nb := range a {
		numCount[nb]++
	}

	for nb, count := range numCount {
		if count%2 != 0 {
			return nb
		}
	}
	return -1
}
